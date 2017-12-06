const async = require("async");
const discovery = require("clever-discovery");
const kayvee = require("kayvee");
const request = require("request");
const opentracing = require("opentracing");
const {commandFactory} = require("hystrixjs");
const RollingNumberEvent = require("hystrixjs/lib/metrics/RollingNumberEvent");

/**
 * @external Span
 * @see {@link https://doc.esdoc.org/github.com/opentracing/opentracing-javascript/class/src/span.js~Span.html}
 */

const { Errors } = require("./types");

/**
 * The exponential retry policy will retry five times with an exponential backoff.
 * @alias module:swagger-test.RetryPolicies.Exponential
 */
const exponentialRetryPolicy = {
  backoffs() {
    const ret = [];
    let next = 100.0; // milliseconds
    const e = 0.05; // +/- 5% jitter
    while (ret.length < 5) {
      const jitter = ((Math.random() * 2) - 1) * e * next;
      ret.push(next + jitter);
      next *= 2;
    }
    return ret;
  },
  retry(requestOptions, err, res) {
    if (err || requestOptions.method === "POST" ||
        requestOptions.method === "PATCH" ||
        res.statusCode < 500) {
      return false;
    }
    return true;
  },
};

/**
 * Use this retry policy to retry a request once.
 * @alias module:swagger-test.RetryPolicies.Single
 */
const singleRetryPolicy = {
  backoffs() {
    return [1000];
  },
  retry(requestOptions, err, res) {
    if (err || requestOptions.method === "POST" ||
        requestOptions.method === "PATCH" ||
        res.statusCode < 500) {
      return false;
    }
    return true;
  },
};

/**
 * Use this retry policy to turn off retries.
 * @alias module:swagger-test.RetryPolicies.None
 */
const noRetryPolicy = {
  backoffs() {
    return [];
  },
  retry() {
    return false;
  },
};

/**
 * Request status log is used to
 * to output the status of a request returned
 * by the client.
 */
function responseLog(logger, req, res, err) {
  var res = res || { };
  var req = req || { };
  var logData = {
	"backend": "swagger-test",
	"method": req.method || "",
	"uri": req.uri || "",
    "message": err || (res.statusMessage || ""),
    "status_code": res.statusCode || 0,
  };

  if (err) {
    logger.errorD("client-request-finished", logData);
  } else {
    logger.infoD("client-request-finished", logData);
  }
}

/**
 * Default circuit breaker options.
 * @alias module:swagger-test.DefaultCircuitOptions
 */
const defaultCircuitOptions = {
  forceClosed:            true,
  requestVolumeThreshold: 20,
  maxConcurrentRequests:  100,
  requestVolumeThreshold: 20,
  sleepWindow:            5000,
  errorPercentThreshold:  90,
  logIntervalMs:          30000
};

/**
 * swagger-test client library.
 * @module swagger-test
 * @typicalname SwaggerTest
 */

/**
 * swagger-test client
 * @alias module:swagger-test
 */
class SwaggerTest {

  /**
   * Create a new client object.
   * @param {Object} options - Options for constructing a client object.
   * @param {string} [options.address] - URL where the server is located. Must provide
   * this or the discovery argument
   * @param {bool} [options.discovery] - Use clever-discovery to locate the server. Must provide
   * this or the address argument
   * @param {number} [options.timeout] - The timeout to use for all client requests,
   * in milliseconds. This can be overridden on a per-request basis.
   * @param {module:swagger-test.RetryPolicies} [options.retryPolicy=RetryPolicies.Single] - The logic to
   * determine which requests to retry, as well as how many times to retry.
   * @param {module:kayvee.Logger} [options.logger=logger.New("swagger-test-wagclient")] - The Kayvee 
   * logger to use in the client.
   * @param {Object} [options.circuit] - Options for constructing the client's circuit breaker.
   * @param {bool} [options.circuit.forceClosed] - When set to true the circuit will always be closed. Default: true.
   * @param {number} [options.circuit.maxConcurrentRequests] - the maximum number of concurrent requests
   * the client can make at the same time. Default: 100.
   * @param {number} [options.circuit.requestVolumeThreshold] - The minimum number of requests needed
   * before a circuit can be tripped due to health. Default: 20.
   * @param {number} [options.circuit.sleepWindow] - how long, in milliseconds, to wait after a circuit opens
   * before testing for recovery. Default: 5000.
   * @param {number} [options.circuit.errorPercentThreshold] - the threshold to place on the rolling error
   * rate. Once the error rate exceeds this percentage, the circuit opens.
   * Default: 90.
   */
  constructor(options) {
    options = options || {};

    if (options.discovery) {
      try {
        this.address = discovery("swagger-test", "http").url();
      } catch (e) {
        this.address = discovery("swagger-test", "default").url();
      }
    } else if (options.address) {
      this.address = options.address;
    } else {
      throw new Error("Cannot initialize swagger-test without discovery or address");
    }
    if (options.timeout) {
      this.timeout = options.timeout;
    }
    if (options.retryPolicy) {
      this.retryPolicy = options.retryPolicy;
    }
    if (options.logger) {
      this.logger = options.logger;
    } else {
      this.logger =  new kayvee.logger("swagger-test-wagclient");
    }

    const circuitOptions = Object.assign({}, defaultCircuitOptions, options.circuit);
    this._hystrixCommand = commandFactory.getOrCreate("swagger-test").
      errorHandler(this._hystrixCommandErrorHandler).
      circuitBreakerForceClosed(circuitOptions.forceClosed).
      requestVolumeRejectionThreshold(circuitOptions.maxConcurrentRequests).
      circuitBreakerRequestVolumeThreshold(circuitOptions.requestVolumeThreshold).
      circuitBreakerSleepWindowInMilliseconds(circuitOptions.sleepWindow).
      circuitBreakerErrorThresholdPercentage(circuitOptions.errorPercentThreshold).
      timeout(0).
      statisticalWindowLength(10000).
      statisticalWindowNumberOfBuckets(10).
      run(this._hystrixCommandRun).
      context(this).
      build();

    setInterval(() => this._logCircuitState(), circuitOptions.logIntervalMs);
  }

  _hystrixCommandErrorHandler(err) {
    // to avoid counting 4XXs as errors, only count an error if it comes from the request library
    if (err._fromRequest === true) {
      return err;
    }
    return false;
  }

  _hystrixCommandRun(method, args) {
    return method.apply(this, args);
  }

  _logCircuitState(logger) {
    // code below heavily borrows from hystrix's internal HystrixSSEStream.js logic
    const metrics = this._hystrixCommand.metrics;
    const healthCounts = metrics.getHealthCounts()
    const circuitBreaker = this._hystrixCommand.circuitBreaker;
    this.logger.infoD("swagger-test", {
      "requestCount":                    healthCounts.totalCount,
      "errorCount":                      healthCounts.errorCount,
      "errorPercentage":                 healthCounts.errorPercentage,
      "isCircuitBreakerOpen":            circuitBreaker.isOpen(),
      "rollingCountFailure":             metrics.getRollingCount(RollingNumberEvent.FAILURE),
      "rollingCountShortCircuited":      metrics.getRollingCount(RollingNumberEvent.SHORT_CIRCUITED),
      "rollingCountSuccess":             metrics.getRollingCount(RollingNumberEvent.SUCCESS),
      "rollingCountTimeout":             metrics.getRollingCount(RollingNumberEvent.TIMEOUT),
      "currentConcurrentExecutionCount": metrics.getCurrentExecutionCount(),
      "latencyTotalMean":                metrics.getExecutionTime("mean") || 0,
    });
  }

  /**
   * Gets authors
   * @param {Object} params
   * @param {string} [params.name]
   * @param {string} [params.startingAfter]
   * @param {object} [options]
   * @param {number} [options.timeout] - A request specific timeout
   * @param {external:Span} [options.span] - An OpenTracing span - For example from the parent request
   * @param {module:swagger-test.RetryPolicies} [options.retryPolicy] - A request specific retryPolicy
   * @param {function} [cb]
   * @returns {Promise}
   * @fulfill {Object}
   * @reject {module:swagger-test.Errors.BadRequest}
   * @reject {module:swagger-test.Errors.InternalError}
   * @reject {Error}
   */
  getAuthors(params, options, cb) {
    return this._hystrixCommand.execute(this._getAuthors, arguments);
  }
  _getAuthors(params, options, cb) {
    if (!cb && typeof options === "function") {
      cb = options;
      options = undefined;
    }

    return new Promise((resolve, reject) => {
      const rejecter = (err) => {
        reject(err);
        if (cb) {
          cb(err);
        }
      };
      const resolver = (data) => {
        resolve(data);
        if (cb) {
          cb(null, data);
        }
      };


      if (!options) {
        options = {};
      }

      const timeout = options.timeout || this.timeout;
      const span = options.span;

      const headers = {};

      const query = {};
      if (typeof params.name !== "undefined") {
        query["name"] = params.name;
      }
  
      if (typeof params.startingAfter !== "undefined") {
        query["startingAfter"] = params.startingAfter;
      }
  

      if (span) {
        opentracing.inject(span, opentracing.FORMAT_TEXT_MAP, headers);
        span.logEvent("GET /v1/authors");
        span.setTag("span.kind", "client");
      }

      const requestOptions = {
        method: "GET",
        uri: this.address + "/v1/authors",
        json: true,
        timeout,
        headers,
        qs: query,
        useQuerystring: true,
      };
  

      const retryPolicy = options.retryPolicy || this.retryPolicy || singleRetryPolicy;
      const backoffs = retryPolicy.backoffs();
      const logger = this.logger;
  
      let retries = 0;
      (function requestOnce() {
        request(requestOptions, (err, response, body) => {
          if (retries < backoffs.length && retryPolicy.retry(requestOptions, err, response, body)) {
            const backoff = backoffs[retries];
            retries += 1;
            setTimeout(requestOnce, backoff);
            return;
          }
          if (err) {
            err._fromRequest = true;
            responseLog(logger, requestOptions, response, err)
            rejecter(err);
            return;
          }

          switch (response.statusCode) {
            case 200:
              resolver(body);
              break;
            
            case 400:
              var err = new Errors.BadRequest(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            case 500:
              var err = new Errors.InternalError(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            default:
              var err = new Error("Received unexpected statusCode " + response.statusCode);
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
          }
        });
      }());
    });
  }


  /**
   * Gets authors
   * @param {Object} params
   * @param {string} [params.name]
   * @param {string} [params.startingAfter]
   * @param {object} [options]
   * @param {number} [options.timeout] - A request specific timeout
   * @param {external:Span} [options.span] - An OpenTracing span - For example from the parent request
   * @param {module:swagger-test.RetryPolicies} [options.retryPolicy] - A request specific retryPolicy
   * @returns {Object} iter
   * @returns {function} iter.map - takes in a function, applies it to each resource, and returns a promise to the result as an array
   * @returns {function} iter.toArray - returns a promise to the resources as an array
   * @returns {function} iter.forEach - takes in a function, applies it to each resource
   */
  getAuthorsIter(params, options) {
    const it = (f, saveResults, cb) => new Promise((resolve, reject) => {
      const rejecter = (err) => {
        reject(err);
        if (cb) {
          cb(err);
        }
      };
      const resolver = (data) => {
        resolve(data);
        if (cb) {
          cb(null, data);
        }
      };


      if (!options) {
        options = {};
      }

      const timeout = options.timeout || this.timeout;
      const span = options.span;

      const headers = {};

      const query = {};
      if (typeof params.name !== "undefined") {
        query["name"] = params.name;
      }
  
      if (typeof params.startingAfter !== "undefined") {
        query["startingAfter"] = params.startingAfter;
      }
  

      if (span) {
        opentracing.inject(span, opentracing.FORMAT_TEXT_MAP, headers);
        span.setTag("span.kind", "client");
      }

      const requestOptions = {
        method: "GET",
        uri: this.address + "/v1/authors",
        json: true,
        timeout,
        headers,
        qs: query,
        useQuerystring: true,
      };
  

      const retryPolicy = options.retryPolicy || this.retryPolicy || singleRetryPolicy;
      const backoffs = retryPolicy.backoffs();
      const logger = this.logger;
  
      let results = [];
      async.whilst(
        () => requestOptions.uri !== "",
        cbW => {
          if (span) {
            span.logEvent("GET /v1/authors");
          }
      const address = this.address;
      let retries = 0;
      (function requestOnce() {
        request(requestOptions, (err, response, body) => {
          if (retries < backoffs.length && retryPolicy.retry(requestOptions, err, response, body)) {
            const backoff = backoffs[retries];
            retries += 1;
            setTimeout(requestOnce, backoff);
            return;
          }
          if (err) {
            err._fromRequest = true;
            responseLog(logger, requestOptions, response, err)
            cbW(err);
            return;
          }

          switch (response.statusCode) {
            case 200:
              if (saveResults) {
                results = results.concat(body.authorSet.results.map(f));
              } else {
                body.authorSet.results.forEach(f);
              }
              break;
            
            case 400:
              var err = new Errors.BadRequest(body || {});
              responseLog(logger, requestOptions, response, err);
              cbW(err);
              return;
            
            case 500:
              var err = new Errors.InternalError(body || {});
              responseLog(logger, requestOptions, response, err);
              cbW(err);
              return;
            
            default:
              var err = new Error("Received unexpected statusCode " + response.statusCode);
              responseLog(logger, requestOptions, response, err);
              cbW(err);
              return;
          }

          requestOptions.qs = null;
          requestOptions.useQuerystring = false;
          requestOptions.uri = "";
          if (response.headers["x-next-page-path"]) {
            requestOptions.uri = address + response.headers["x-next-page-path"];
          }
          cbW();
        });
      }());
        },
        err => {
          if (err) {
            rejecter(err);
            return;
          }
          if (saveResults) {
            resolver(results);
          } else {
            resolver();
          }
        }
      );
    });

    return {
      map: (f, cb) => this._hystrixCommand.execute(it, [f, true, cb]),
      toArray: cb => this._hystrixCommand.execute(it, [x => x, true, cb]),
      forEach: (f, cb) => this._hystrixCommand.execute(it, [f, false, cb]),
    };
  }

  /**
   * Gets authors, but needs to use the body so it's a PUT
   * @param {Object} params
   * @param {string} [params.name]
   * @param {string} [params.startingAfter]
   * @param [params.favoriteBooks]
   * @param {object} [options]
   * @param {number} [options.timeout] - A request specific timeout
   * @param {external:Span} [options.span] - An OpenTracing span - For example from the parent request
   * @param {module:swagger-test.RetryPolicies} [options.retryPolicy] - A request specific retryPolicy
   * @param {function} [cb]
   * @returns {Promise}
   * @fulfill {Object}
   * @reject {module:swagger-test.Errors.BadRequest}
   * @reject {module:swagger-test.Errors.InternalError}
   * @reject {Error}
   */
  getAuthorsWithPut(params, options, cb) {
    return this._hystrixCommand.execute(this._getAuthorsWithPut, arguments);
  }
  _getAuthorsWithPut(params, options, cb) {
    if (!cb && typeof options === "function") {
      cb = options;
      options = undefined;
    }

    return new Promise((resolve, reject) => {
      const rejecter = (err) => {
        reject(err);
        if (cb) {
          cb(err);
        }
      };
      const resolver = (data) => {
        resolve(data);
        if (cb) {
          cb(null, data);
        }
      };


      if (!options) {
        options = {};
      }

      const timeout = options.timeout || this.timeout;
      const span = options.span;

      const headers = {};

      const query = {};
      if (typeof params.name !== "undefined") {
        query["name"] = params.name;
      }
  
      if (typeof params.startingAfter !== "undefined") {
        query["startingAfter"] = params.startingAfter;
      }
  

      if (span) {
        opentracing.inject(span, opentracing.FORMAT_TEXT_MAP, headers);
        span.logEvent("PUT /v1/authors");
        span.setTag("span.kind", "client");
      }

      const requestOptions = {
        method: "PUT",
        uri: this.address + "/v1/authors",
        json: true,
        timeout,
        headers,
        qs: query,
        useQuerystring: true,
      };
  
      requestOptions.body = params.favoriteBooks;
  

      const retryPolicy = options.retryPolicy || this.retryPolicy || singleRetryPolicy;
      const backoffs = retryPolicy.backoffs();
      const logger = this.logger;
  
      let retries = 0;
      (function requestOnce() {
        request(requestOptions, (err, response, body) => {
          if (retries < backoffs.length && retryPolicy.retry(requestOptions, err, response, body)) {
            const backoff = backoffs[retries];
            retries += 1;
            setTimeout(requestOnce, backoff);
            return;
          }
          if (err) {
            err._fromRequest = true;
            responseLog(logger, requestOptions, response, err)
            rejecter(err);
            return;
          }

          switch (response.statusCode) {
            case 200:
              resolver(body);
              break;
            
            case 400:
              var err = new Errors.BadRequest(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            case 500:
              var err = new Errors.InternalError(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            default:
              var err = new Error("Received unexpected statusCode " + response.statusCode);
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
          }
        });
      }());
    });
  }


  /**
   * Gets authors, but needs to use the body so it's a PUT
   * @param {Object} params
   * @param {string} [params.name]
   * @param {string} [params.startingAfter]
   * @param [params.favoriteBooks]
   * @param {object} [options]
   * @param {number} [options.timeout] - A request specific timeout
   * @param {external:Span} [options.span] - An OpenTracing span - For example from the parent request
   * @param {module:swagger-test.RetryPolicies} [options.retryPolicy] - A request specific retryPolicy
   * @returns {Object} iter
   * @returns {function} iter.map - takes in a function, applies it to each resource, and returns a promise to the result as an array
   * @returns {function} iter.toArray - returns a promise to the resources as an array
   * @returns {function} iter.forEach - takes in a function, applies it to each resource
   */
  getAuthorsWithPutIter(params, options) {
    const it = (f, saveResults, cb) => new Promise((resolve, reject) => {
      const rejecter = (err) => {
        reject(err);
        if (cb) {
          cb(err);
        }
      };
      const resolver = (data) => {
        resolve(data);
        if (cb) {
          cb(null, data);
        }
      };


      if (!options) {
        options = {};
      }

      const timeout = options.timeout || this.timeout;
      const span = options.span;

      const headers = {};

      const query = {};
      if (typeof params.name !== "undefined") {
        query["name"] = params.name;
      }
  
      if (typeof params.startingAfter !== "undefined") {
        query["startingAfter"] = params.startingAfter;
      }
  

      if (span) {
        opentracing.inject(span, opentracing.FORMAT_TEXT_MAP, headers);
        span.setTag("span.kind", "client");
      }

      const requestOptions = {
        method: "PUT",
        uri: this.address + "/v1/authors",
        json: true,
        timeout,
        headers,
        qs: query,
        useQuerystring: true,
      };
  
      requestOptions.body = params.favoriteBooks;
  

      const retryPolicy = options.retryPolicy || this.retryPolicy || singleRetryPolicy;
      const backoffs = retryPolicy.backoffs();
      const logger = this.logger;
  
      let results = [];
      async.whilst(
        () => requestOptions.uri !== "",
        cbW => {
          if (span) {
            span.logEvent("PUT /v1/authors");
          }
      const address = this.address;
      let retries = 0;
      (function requestOnce() {
        request(requestOptions, (err, response, body) => {
          if (retries < backoffs.length && retryPolicy.retry(requestOptions, err, response, body)) {
            const backoff = backoffs[retries];
            retries += 1;
            setTimeout(requestOnce, backoff);
            return;
          }
          if (err) {
            err._fromRequest = true;
            responseLog(logger, requestOptions, response, err)
            cbW(err);
            return;
          }

          switch (response.statusCode) {
            case 200:
              if (saveResults) {
                results = results.concat(body.authorSet.results.map(f));
              } else {
                body.authorSet.results.forEach(f);
              }
              break;
            
            case 400:
              var err = new Errors.BadRequest(body || {});
              responseLog(logger, requestOptions, response, err);
              cbW(err);
              return;
            
            case 500:
              var err = new Errors.InternalError(body || {});
              responseLog(logger, requestOptions, response, err);
              cbW(err);
              return;
            
            default:
              var err = new Error("Received unexpected statusCode " + response.statusCode);
              responseLog(logger, requestOptions, response, err);
              cbW(err);
              return;
          }

          requestOptions.qs = null;
          requestOptions.useQuerystring = false;
          requestOptions.uri = "";
          if (response.headers["x-next-page-path"]) {
            requestOptions.uri = address + response.headers["x-next-page-path"];
          }
          cbW();
        });
      }());
        },
        err => {
          if (err) {
            rejecter(err);
            return;
          }
          if (saveResults) {
            resolver(results);
          } else {
            resolver();
          }
        }
      );
    });

    return {
      map: (f, cb) => this._hystrixCommand.execute(it, [f, true, cb]),
      toArray: cb => this._hystrixCommand.execute(it, [x => x, true, cb]),
      forEach: (f, cb) => this._hystrixCommand.execute(it, [f, false, cb]),
    };
  }

  /**
   * Returns a list of books
   * @param {Object} params
   * @param {string[]} [params.authors] - A list of authors. Must specify at least one and at most two
   * @param {boolean} [params.available=true]
   * @param {string} [params.state=finished]
   * @param {string} [params.published]
   * @param {string} [params.snakeCase]
   * @param {string} [params.completed]
   * @param {number} [params.maxPages=500.5]
   * @param {number} [params.minPages=5]
   * @param {number} [params.pagesToTime]
   * @param {string} [params.authorization]
   * @param {number} [params.startingAfter]
   * @param {object} [options]
   * @param {number} [options.timeout] - A request specific timeout
   * @param {external:Span} [options.span] - An OpenTracing span - For example from the parent request
   * @param {module:swagger-test.RetryPolicies} [options.retryPolicy] - A request specific retryPolicy
   * @param {function} [cb]
   * @returns {Promise}
   * @fulfill {Object[]}
   * @reject {module:swagger-test.Errors.BadRequest}
   * @reject {module:swagger-test.Errors.InternalError}
   * @reject {Error}
   */
  getBooks(params, options, cb) {
    return this._hystrixCommand.execute(this._getBooks, arguments);
  }
  _getBooks(params, options, cb) {
    if (!cb && typeof options === "function") {
      cb = options;
      options = undefined;
    }

    return new Promise((resolve, reject) => {
      const rejecter = (err) => {
        reject(err);
        if (cb) {
          cb(err);
        }
      };
      const resolver = (data) => {
        resolve(data);
        if (cb) {
          cb(null, data);
        }
      };


      if (!options) {
        options = {};
      }

      const timeout = options.timeout || this.timeout;
      const span = options.span;

      const headers = {};
      headers["authorization"] = params.authorization;

      const query = {};
      if (typeof params.authors !== "undefined") {
        query["authors"] = params.authors;
      }
  
      if (typeof params.available !== "undefined") {
        query["available"] = params.available;
      }
  
      if (typeof params.state !== "undefined") {
        query["state"] = params.state;
      }
  
      if (typeof params.published !== "undefined") {
        query["published"] = params.published;
      }
  
      if (typeof params.snakeCase !== "undefined") {
        query["snake_case"] = params.snakeCase;
      }
  
      if (typeof params.completed !== "undefined") {
        query["completed"] = params.completed;
      }
  
      if (typeof params.maxPages !== "undefined") {
        query["maxPages"] = params.maxPages;
      }
  
      if (typeof params.minPages !== "undefined") {
        query["min_pages"] = params.minPages;
      }
  
      if (typeof params.pagesToTime !== "undefined") {
        query["pagesToTime"] = params.pagesToTime;
      }
  
      if (typeof params.startingAfter !== "undefined") {
        query["startingAfter"] = params.startingAfter;
      }
  

      if (span) {
        opentracing.inject(span, opentracing.FORMAT_TEXT_MAP, headers);
        span.logEvent("GET /v1/books");
        span.setTag("span.kind", "client");
      }

      const requestOptions = {
        method: "GET",
        uri: this.address + "/v1/books",
        json: true,
        timeout,
        headers,
        qs: query,
        useQuerystring: true,
      };
  

      const retryPolicy = options.retryPolicy || this.retryPolicy || singleRetryPolicy;
      const backoffs = retryPolicy.backoffs();
      const logger = this.logger;
  
      let retries = 0;
      (function requestOnce() {
        request(requestOptions, (err, response, body) => {
          if (retries < backoffs.length && retryPolicy.retry(requestOptions, err, response, body)) {
            const backoff = backoffs[retries];
            retries += 1;
            setTimeout(requestOnce, backoff);
            return;
          }
          if (err) {
            err._fromRequest = true;
            responseLog(logger, requestOptions, response, err)
            rejecter(err);
            return;
          }

          switch (response.statusCode) {
            case 200:
              resolver(body);
              break;
            
            case 400:
              var err = new Errors.BadRequest(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            case 500:
              var err = new Errors.InternalError(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            default:
              var err = new Error("Received unexpected statusCode " + response.statusCode);
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
          }
        });
      }());
    });
  }


  /**
   * Returns a list of books
   * @param {Object} params
   * @param {string[]} [params.authors] - A list of authors. Must specify at least one and at most two
   * @param {boolean} [params.available=true]
   * @param {string} [params.state=finished]
   * @param {string} [params.published]
   * @param {string} [params.snakeCase]
   * @param {string} [params.completed]
   * @param {number} [params.maxPages=500.5]
   * @param {number} [params.minPages=5]
   * @param {number} [params.pagesToTime]
   * @param {string} [params.authorization]
   * @param {number} [params.startingAfter]
   * @param {object} [options]
   * @param {number} [options.timeout] - A request specific timeout
   * @param {external:Span} [options.span] - An OpenTracing span - For example from the parent request
   * @param {module:swagger-test.RetryPolicies} [options.retryPolicy] - A request specific retryPolicy
   * @returns {Object} iter
   * @returns {function} iter.map - takes in a function, applies it to each resource, and returns a promise to the result as an array
   * @returns {function} iter.toArray - returns a promise to the resources as an array
   * @returns {function} iter.forEach - takes in a function, applies it to each resource
   */
  getBooksIter(params, options) {
    const it = (f, saveResults, cb) => new Promise((resolve, reject) => {
      const rejecter = (err) => {
        reject(err);
        if (cb) {
          cb(err);
        }
      };
      const resolver = (data) => {
        resolve(data);
        if (cb) {
          cb(null, data);
        }
      };


      if (!options) {
        options = {};
      }

      const timeout = options.timeout || this.timeout;
      const span = options.span;

      const headers = {};
      headers["authorization"] = params.authorization;

      const query = {};
      if (typeof params.authors !== "undefined") {
        query["authors"] = params.authors;
      }
  
      if (typeof params.available !== "undefined") {
        query["available"] = params.available;
      }
  
      if (typeof params.state !== "undefined") {
        query["state"] = params.state;
      }
  
      if (typeof params.published !== "undefined") {
        query["published"] = params.published;
      }
  
      if (typeof params.snakeCase !== "undefined") {
        query["snake_case"] = params.snakeCase;
      }
  
      if (typeof params.completed !== "undefined") {
        query["completed"] = params.completed;
      }
  
      if (typeof params.maxPages !== "undefined") {
        query["maxPages"] = params.maxPages;
      }
  
      if (typeof params.minPages !== "undefined") {
        query["min_pages"] = params.minPages;
      }
  
      if (typeof params.pagesToTime !== "undefined") {
        query["pagesToTime"] = params.pagesToTime;
      }
  
      if (typeof params.startingAfter !== "undefined") {
        query["startingAfter"] = params.startingAfter;
      }
  

      if (span) {
        opentracing.inject(span, opentracing.FORMAT_TEXT_MAP, headers);
        span.setTag("span.kind", "client");
      }

      const requestOptions = {
        method: "GET",
        uri: this.address + "/v1/books",
        json: true,
        timeout,
        headers,
        qs: query,
        useQuerystring: true,
      };
  

      const retryPolicy = options.retryPolicy || this.retryPolicy || singleRetryPolicy;
      const backoffs = retryPolicy.backoffs();
      const logger = this.logger;
  
      let results = [];
      async.whilst(
        () => requestOptions.uri !== "",
        cbW => {
          if (span) {
            span.logEvent("GET /v1/books");
          }
      const address = this.address;
      let retries = 0;
      (function requestOnce() {
        request(requestOptions, (err, response, body) => {
          if (retries < backoffs.length && retryPolicy.retry(requestOptions, err, response, body)) {
            const backoff = backoffs[retries];
            retries += 1;
            setTimeout(requestOnce, backoff);
            return;
          }
          if (err) {
            err._fromRequest = true;
            responseLog(logger, requestOptions, response, err)
            cbW(err);
            return;
          }

          switch (response.statusCode) {
            case 200:
              if (saveResults) {
                results = results.concat(body.map(f));
              } else {
                body.forEach(f);
              }
              break;
            
            case 400:
              var err = new Errors.BadRequest(body || {});
              responseLog(logger, requestOptions, response, err);
              cbW(err);
              return;
            
            case 500:
              var err = new Errors.InternalError(body || {});
              responseLog(logger, requestOptions, response, err);
              cbW(err);
              return;
            
            default:
              var err = new Error("Received unexpected statusCode " + response.statusCode);
              responseLog(logger, requestOptions, response, err);
              cbW(err);
              return;
          }

          requestOptions.qs = null;
          requestOptions.useQuerystring = false;
          requestOptions.uri = "";
          if (response.headers["x-next-page-path"]) {
            requestOptions.uri = address + response.headers["x-next-page-path"];
          }
          cbW();
        });
      }());
        },
        err => {
          if (err) {
            rejecter(err);
            return;
          }
          if (saveResults) {
            resolver(results);
          } else {
            resolver();
          }
        }
      );
    });

    return {
      map: (f, cb) => this._hystrixCommand.execute(it, [f, true, cb]),
      toArray: cb => this._hystrixCommand.execute(it, [x => x, true, cb]),
      forEach: (f, cb) => this._hystrixCommand.execute(it, [f, false, cb]),
    };
  }

  /**
   * Creates a book
   * @param newBook
   * @param {object} [options]
   * @param {number} [options.timeout] - A request specific timeout
   * @param {external:Span} [options.span] - An OpenTracing span - For example from the parent request
   * @param {module:swagger-test.RetryPolicies} [options.retryPolicy] - A request specific retryPolicy
   * @param {function} [cb]
   * @returns {Promise}
   * @fulfill {Object}
   * @reject {module:swagger-test.Errors.BadRequest}
   * @reject {module:swagger-test.Errors.InternalError}
   * @reject {Error}
   */
  createBook(newBook, options, cb) {
    return this._hystrixCommand.execute(this._createBook, arguments);
  }
  _createBook(newBook, options, cb) {
    const params = {};
    params["newBook"] = newBook;

    if (!cb && typeof options === "function") {
      cb = options;
      options = undefined;
    }

    return new Promise((resolve, reject) => {
      const rejecter = (err) => {
        reject(err);
        if (cb) {
          cb(err);
        }
      };
      const resolver = (data) => {
        resolve(data);
        if (cb) {
          cb(null, data);
        }
      };


      if (!options) {
        options = {};
      }

      const timeout = options.timeout || this.timeout;
      const span = options.span;

      const headers = {};

      const query = {};

      if (span) {
        opentracing.inject(span, opentracing.FORMAT_TEXT_MAP, headers);
        span.logEvent("POST /v1/books");
        span.setTag("span.kind", "client");
      }

      const requestOptions = {
        method: "POST",
        uri: this.address + "/v1/books",
        json: true,
        timeout,
        headers,
        qs: query,
        useQuerystring: true,
      };
  
      requestOptions.body = params.newBook;
  

      const retryPolicy = options.retryPolicy || this.retryPolicy || singleRetryPolicy;
      const backoffs = retryPolicy.backoffs();
      const logger = this.logger;
  
      let retries = 0;
      (function requestOnce() {
        request(requestOptions, (err, response, body) => {
          if (retries < backoffs.length && retryPolicy.retry(requestOptions, err, response, body)) {
            const backoff = backoffs[retries];
            retries += 1;
            setTimeout(requestOnce, backoff);
            return;
          }
          if (err) {
            err._fromRequest = true;
            responseLog(logger, requestOptions, response, err)
            rejecter(err);
            return;
          }

          switch (response.statusCode) {
            case 200:
              resolver(body);
              break;
            
            case 400:
              var err = new Errors.BadRequest(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            case 500:
              var err = new Errors.InternalError(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            default:
              var err = new Error("Received unexpected statusCode " + response.statusCode);
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
          }
        });
      }());
    });
  }

  /**
   * Puts a book
   * @param newBook
   * @param {object} [options]
   * @param {number} [options.timeout] - A request specific timeout
   * @param {external:Span} [options.span] - An OpenTracing span - For example from the parent request
   * @param {module:swagger-test.RetryPolicies} [options.retryPolicy] - A request specific retryPolicy
   * @param {function} [cb]
   * @returns {Promise}
   * @fulfill {Object}
   * @reject {module:swagger-test.Errors.BadRequest}
   * @reject {module:swagger-test.Errors.InternalError}
   * @reject {Error}
   */
  putBook(newBook, options, cb) {
    return this._hystrixCommand.execute(this._putBook, arguments);
  }
  _putBook(newBook, options, cb) {
    const params = {};
    params["newBook"] = newBook;

    if (!cb && typeof options === "function") {
      cb = options;
      options = undefined;
    }

    return new Promise((resolve, reject) => {
      const rejecter = (err) => {
        reject(err);
        if (cb) {
          cb(err);
        }
      };
      const resolver = (data) => {
        resolve(data);
        if (cb) {
          cb(null, data);
        }
      };


      if (!options) {
        options = {};
      }

      const timeout = options.timeout || this.timeout;
      const span = options.span;

      const headers = {};

      const query = {};

      if (span) {
        opentracing.inject(span, opentracing.FORMAT_TEXT_MAP, headers);
        span.logEvent("PUT /v1/books");
        span.setTag("span.kind", "client");
      }

      const requestOptions = {
        method: "PUT",
        uri: this.address + "/v1/books",
        json: true,
        timeout,
        headers,
        qs: query,
        useQuerystring: true,
      };
  
      requestOptions.body = params.newBook;
  

      const retryPolicy = options.retryPolicy || this.retryPolicy || singleRetryPolicy;
      const backoffs = retryPolicy.backoffs();
      const logger = this.logger;
  
      let retries = 0;
      (function requestOnce() {
        request(requestOptions, (err, response, body) => {
          if (retries < backoffs.length && retryPolicy.retry(requestOptions, err, response, body)) {
            const backoff = backoffs[retries];
            retries += 1;
            setTimeout(requestOnce, backoff);
            return;
          }
          if (err) {
            err._fromRequest = true;
            responseLog(logger, requestOptions, response, err)
            rejecter(err);
            return;
          }

          switch (response.statusCode) {
            case 200:
              resolver(body);
              break;
            
            case 400:
              var err = new Errors.BadRequest(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            case 500:
              var err = new Errors.InternalError(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            default:
              var err = new Error("Received unexpected statusCode " + response.statusCode);
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
          }
        });
      }());
    });
  }

  /**
   * Returns a book
   * @param {Object} params
   * @param {number} params.bookID
   * @param {string} [params.authorID]
   * @param {string} [params.authorization]
   * @param {string} [params.XDontRateLimitMeBro]
   * @param {string} [params.randomBytes]
   * @param {object} [options]
   * @param {number} [options.timeout] - A request specific timeout
   * @param {external:Span} [options.span] - An OpenTracing span - For example from the parent request
   * @param {module:swagger-test.RetryPolicies} [options.retryPolicy] - A request specific retryPolicy
   * @param {function} [cb]
   * @returns {Promise}
   * @fulfill {Object}
   * @reject {module:swagger-test.Errors.BadRequest}
   * @reject {module:swagger-test.Errors.Unathorized}
   * @reject {module:swagger-test.Errors.Error}
   * @reject {module:swagger-test.Errors.InternalError}
   * @reject {Error}
   */
  getBookByID(params, options, cb) {
    return this._hystrixCommand.execute(this._getBookByID, arguments);
  }
  _getBookByID(params, options, cb) {
    if (!cb && typeof options === "function") {
      cb = options;
      options = undefined;
    }

    return new Promise((resolve, reject) => {
      const rejecter = (err) => {
        reject(err);
        if (cb) {
          cb(err);
        }
      };
      const resolver = (data) => {
        resolve(data);
        if (cb) {
          cb(null, data);
        }
      };


      if (!options) {
        options = {};
      }

      const timeout = options.timeout || this.timeout;
      const span = options.span;

      const headers = {};
      if (!params.bookID) {
        rejecter(new Error("bookID must be non-empty because it's a path parameter"));
        return;
      }
      headers["authorization"] = params.authorization;
      headers["X-Dont-Rate-Limit-Me-Bro"] = params.XDontRateLimitMeBro;

      const query = {};
      if (typeof params.authorID !== "undefined") {
        query["authorID"] = params.authorID;
      }
  
      if (typeof params.randomBytes !== "undefined") {
        query["randomBytes"] = params.randomBytes;
      }
  

      if (span) {
        opentracing.inject(span, opentracing.FORMAT_TEXT_MAP, headers);
        span.logEvent("GET /v1/books/{book_id}");
        span.setTag("span.kind", "client");
      }

      const requestOptions = {
        method: "GET",
        uri: this.address + "/v1/books/" + params.bookID + "",
        json: true,
        timeout,
        headers,
        qs: query,
        useQuerystring: true,
      };
  

      const retryPolicy = options.retryPolicy || this.retryPolicy || singleRetryPolicy;
      const backoffs = retryPolicy.backoffs();
      const logger = this.logger;
  
      let retries = 0;
      (function requestOnce() {
        request(requestOptions, (err, response, body) => {
          if (retries < backoffs.length && retryPolicy.retry(requestOptions, err, response, body)) {
            const backoff = backoffs[retries];
            retries += 1;
            setTimeout(requestOnce, backoff);
            return;
          }
          if (err) {
            err._fromRequest = true;
            responseLog(logger, requestOptions, response, err)
            rejecter(err);
            return;
          }

          switch (response.statusCode) {
            case 200:
              resolver(body);
              break;
            
            case 400:
              var err = new Errors.BadRequest(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            case 401:
              var err = new Errors.Unathorized(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            case 404:
              var err = new Errors.Error(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            case 500:
              var err = new Errors.InternalError(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            default:
              var err = new Error("Received unexpected statusCode " + response.statusCode);
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
          }
        });
      }());
    });
  }

  /**
   * Retrieve a book
   * @param {string} id
   * @param {object} [options]
   * @param {number} [options.timeout] - A request specific timeout
   * @param {external:Span} [options.span] - An OpenTracing span - For example from the parent request
   * @param {module:swagger-test.RetryPolicies} [options.retryPolicy] - A request specific retryPolicy
   * @param {function} [cb]
   * @returns {Promise}
   * @fulfill {Object}
   * @reject {module:swagger-test.Errors.BadRequest}
   * @reject {module:swagger-test.Errors.Error}
   * @reject {module:swagger-test.Errors.InternalError}
   * @reject {Error}
   */
  getBookByID2(id, options, cb) {
    return this._hystrixCommand.execute(this._getBookByID2, arguments);
  }
  _getBookByID2(id, options, cb) {
    const params = {};
    params["id"] = id;

    if (!cb && typeof options === "function") {
      cb = options;
      options = undefined;
    }

    return new Promise((resolve, reject) => {
      const rejecter = (err) => {
        reject(err);
        if (cb) {
          cb(err);
        }
      };
      const resolver = (data) => {
        resolve(data);
        if (cb) {
          cb(null, data);
        }
      };


      if (!options) {
        options = {};
      }

      const timeout = options.timeout || this.timeout;
      const span = options.span;

      const headers = {};
      if (!params.id) {
        rejecter(new Error("id must be non-empty because it's a path parameter"));
        return;
      }

      const query = {};

      if (span) {
        opentracing.inject(span, opentracing.FORMAT_TEXT_MAP, headers);
        span.logEvent("GET /v1/books2/{id}");
        span.setTag("span.kind", "client");
      }

      const requestOptions = {
        method: "GET",
        uri: this.address + "/v1/books2/" + params.id + "",
        json: true,
        timeout,
        headers,
        qs: query,
        useQuerystring: true,
      };
  

      const retryPolicy = options.retryPolicy || this.retryPolicy || singleRetryPolicy;
      const backoffs = retryPolicy.backoffs();
      const logger = this.logger;
  
      let retries = 0;
      (function requestOnce() {
        request(requestOptions, (err, response, body) => {
          if (retries < backoffs.length && retryPolicy.retry(requestOptions, err, response, body)) {
            const backoff = backoffs[retries];
            retries += 1;
            setTimeout(requestOnce, backoff);
            return;
          }
          if (err) {
            err._fromRequest = true;
            responseLog(logger, requestOptions, response, err)
            rejecter(err);
            return;
          }

          switch (response.statusCode) {
            case 200:
              resolver(body);
              break;
            
            case 400:
              var err = new Errors.BadRequest(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            case 404:
              var err = new Errors.Error(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            case 500:
              var err = new Errors.InternalError(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            default:
              var err = new Error("Received unexpected statusCode " + response.statusCode);
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
          }
        });
      }());
    });
  }

  /**
   * Retrieve a book
   * @param {string} id
   * @param {object} [options]
   * @param {number} [options.timeout] - A request specific timeout
   * @param {external:Span} [options.span] - An OpenTracing span - For example from the parent request
   * @param {module:swagger-test.RetryPolicies} [options.retryPolicy] - A request specific retryPolicy
   * @param {function} [cb]
   * @returns {Promise}
   * @fulfill {Object}
   * @reject {module:swagger-test.Errors.BadRequest}
   * @reject {module:swagger-test.Errors.Error}
   * @reject {module:swagger-test.Errors.InternalError}
   * @reject {Error}
   */
  getBookByIDCached(id, options, cb) {
    return this._hystrixCommand.execute(this._getBookByIDCached, arguments);
  }
  _getBookByIDCached(id, options, cb) {
    const params = {};
    params["id"] = id;

    if (!cb && typeof options === "function") {
      cb = options;
      options = undefined;
    }

    return new Promise((resolve, reject) => {
      const rejecter = (err) => {
        reject(err);
        if (cb) {
          cb(err);
        }
      };
      const resolver = (data) => {
        resolve(data);
        if (cb) {
          cb(null, data);
        }
      };


      if (!options) {
        options = {};
      }

      const timeout = options.timeout || this.timeout;
      const span = options.span;

      const headers = {};
      if (!params.id) {
        rejecter(new Error("id must be non-empty because it's a path parameter"));
        return;
      }

      const query = {};

      if (span) {
        opentracing.inject(span, opentracing.FORMAT_TEXT_MAP, headers);
        span.logEvent("GET /v1/bookscached/{id}");
        span.setTag("span.kind", "client");
      }

      const requestOptions = {
        method: "GET",
        uri: this.address + "/v1/bookscached/" + params.id + "",
        json: true,
        timeout,
        headers,
        qs: query,
        useQuerystring: true,
      };
  

      const retryPolicy = options.retryPolicy || this.retryPolicy || singleRetryPolicy;
      const backoffs = retryPolicy.backoffs();
      const logger = this.logger;
  
      let retries = 0;
      (function requestOnce() {
        request(requestOptions, (err, response, body) => {
          if (retries < backoffs.length && retryPolicy.retry(requestOptions, err, response, body)) {
            const backoff = backoffs[retries];
            retries += 1;
            setTimeout(requestOnce, backoff);
            return;
          }
          if (err) {
            err._fromRequest = true;
            responseLog(logger, requestOptions, response, err)
            rejecter(err);
            return;
          }

          switch (response.statusCode) {
            case 200:
              resolver(body);
              break;
            
            case 400:
              var err = new Errors.BadRequest(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            case 404:
              var err = new Errors.Error(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            case 500:
              var err = new Errors.InternalError(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            default:
              var err = new Error("Received unexpected statusCode " + response.statusCode);
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
          }
        });
      }());
    });
  }

  /**
   * @param {object} [options]
   * @param {number} [options.timeout] - A request specific timeout
   * @param {external:Span} [options.span] - An OpenTracing span - For example from the parent request
   * @param {module:swagger-test.RetryPolicies} [options.retryPolicy] - A request specific retryPolicy
   * @param {function} [cb]
   * @returns {Promise}
   * @fulfill {undefined}
   * @reject {module:swagger-test.Errors.BadRequest}
   * @reject {module:swagger-test.Errors.InternalError}
   * @reject {Error}
   */
  healthCheck(options, cb) {
    return this._hystrixCommand.execute(this._healthCheck, arguments);
  }
  _healthCheck(options, cb) {
    const params = {};

    if (!cb && typeof options === "function") {
      cb = options;
      options = undefined;
    }

    return new Promise((resolve, reject) => {
      const rejecter = (err) => {
        reject(err);
        if (cb) {
          cb(err);
        }
      };
      const resolver = (data) => {
        resolve(data);
        if (cb) {
          cb(null, data);
        }
      };


      if (!options) {
        options = {};
      }

      const timeout = options.timeout || this.timeout;
      const span = options.span;

      const headers = {};

      const query = {};

      if (span) {
        opentracing.inject(span, opentracing.FORMAT_TEXT_MAP, headers);
        span.logEvent("GET /v1/health/check");
        span.setTag("span.kind", "client");
      }

      const requestOptions = {
        method: "GET",
        uri: this.address + "/v1/health/check",
        json: true,
        timeout,
        headers,
        qs: query,
        useQuerystring: true,
      };
  

      const retryPolicy = options.retryPolicy || this.retryPolicy || singleRetryPolicy;
      const backoffs = retryPolicy.backoffs();
      const logger = this.logger;
  
      let retries = 0;
      (function requestOnce() {
        request(requestOptions, (err, response, body) => {
          if (retries < backoffs.length && retryPolicy.retry(requestOptions, err, response, body)) {
            const backoff = backoffs[retries];
            retries += 1;
            setTimeout(requestOnce, backoff);
            return;
          }
          if (err) {
            err._fromRequest = true;
            responseLog(logger, requestOptions, response, err)
            rejecter(err);
            return;
          }

          switch (response.statusCode) {
            case 200:
              resolver();
              break;
            
            case 400:
              var err = new Errors.BadRequest(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            case 500:
              var err = new Errors.InternalError(body || {});
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
            
            default:
              var err = new Error("Received unexpected statusCode " + response.statusCode);
              responseLog(logger, requestOptions, response, err);
              rejecter(err);
              return;
          }
        });
      }());
    });
  }
};

module.exports = SwaggerTest;

/**
 * Retry policies available to use.
 * @alias module:swagger-test.RetryPolicies
 */
module.exports.RetryPolicies = {
  Single: singleRetryPolicy,
  Exponential: exponentialRetryPolicy,
  None: noRetryPolicy,
};

/**
 * Errors returned by methods.
 * @alias module:swagger-test.Errors
 */
module.exports.Errors = Errors;

module.exports.DefaultCircuitOptions = defaultCircuitOptions;
