const discovery = require("@clever/discovery");
const request = require("request");
const url = require("url");
const opentracing = require("opentracing");

// go-swagger treats handles/expects arrays in the query string to be a string of comma joined values
// so...do that thing. It's worth noting that this has lots of issues ("what if my values have commas in them?")
// but that's an issue with go-swagger
function serializeQueryString(data) {
  if (Array.isArray(data)) {
    return data.join(",");
  }
  return data;
}

<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> Sort responses
const defaultRetryPolicy = {
  backoffs() {
    const ret = [];
    let next = 100.0; // milliseconds
    const e = 0.05; // +/- 5% jitter
    while (ret.length < 5) {
      const jitter = (Math.random()*2-1.0)*e*next;
      ret.push(next + jitter);
      next *= 2;
    }
    return ret;
  },
  retry(requestOptions, err, res, body) {
    if (err || requestOptions.method === "POST" ||
        requestOptions.method === "PATCH" ||
        res.statusCode < 500) {
      return false;
    }
    return true;
  },
};

const noRetryPolicy = {
  backoffs() {
    return [];
  },
  retry(requestOptions, err, res, body) {
    return false;
  },
};

<<<<<<< HEAD
=======
>>>>>>> Generated
=======
>>>>>>> Sort responses
module.exports = class SwaggerTest {

  constructor(options) {
    options = options || {};

    if (options.discovery) {
      try {
        this.address = discovery("swagger-test", "http").url();
      } catch (e) {
        this.address = discovery("swagger-test", "default").url();
      };
    } else if (options.address) {
      this.address = options.address;
    } else {
      throw new Error("Cannot initialize swagger-test without discovery or address");
    }
    if (options.timeout) {
      this.timeout = options.timeout
    }
<<<<<<< HEAD
<<<<<<< HEAD
    if (options.retryPolicy) {
      this.retryPolicy = options.retryPolicy;
    }
=======
>>>>>>> Generated
=======
    if (options.retryPolicy) {
      this.retryPolicy = options.retryPolicy;
    }
>>>>>>> Sort responses
  }

  getBook(id, options, cb) {
    const params = {};
    params["id"] = id;

    if (!cb && typeof options === "function") {
      cb = options;
      options = undefined;
    }

    if (!options) {
      options = {};
    }

    const timeout = options.timeout || this.timeout;
    const span = options.span;

    const headers = {};

    const query = {};

    if (span) {
      opentracing.inject(span, opentracing.FORMAT_TEXT_MAP, headers);
      span.logEvent("GET /v1/books/{id}");
    }

    const requestOptions = {
      method: "GET",
      uri: this.address + "/v1/books/" + params.id + "",
      json: true,
      timeout,
      headers,
      qs: query,
    };

    return new Promise((resolve, reject) => {
      const rejecter = (err) => {
        reject(err);
        if (cb) {
          cb(err);
        }
      }
      const resolver = (data) => {
        resolve(data);
        if (cb) {
          cb(null, data);
        }
      }

<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> Sort responses
      const retryPolicy = options.retryPolicy || this.retryPolicy || defaultRetryPolicy;
      const backoffs = retryPolicy.backoffs();
      let retries = 0;
      (function requestOnce() {
        request(requestOptions, (err, response, body) => {
          if (retries < backoffs.length && retryPolicy.retry(requestOptions, err, response, body)) {
            const backoff = backoffs[retries];
            retries += 1;
            return setTimeout(requestOnce, backoff);
          }
          if (err) {
            return rejecter(err);
          }
          if (response.statusCode >= 400) {
            return rejecter(new Error(body));
          }
          resolver(body);
        });
      })();
<<<<<<< HEAD
    });
  }
}

module.exports.RetryPolicies = {
  Default: defaultRetryPolicy,
  None: noRetryPolicy,
};
=======
      request(requestOptions, (err, response, body) => {
        if (err) {
          return rejecter(err);
        }
        if (response.statusCode >= 400) {
          return rejecter(new Error(body));
        }
        resolver(body);
      });
    });
  }
}
>>>>>>> Generated
=======
    });
  }
}

module.exports.RetryPolicies = {
  Default: defaultRetryPolicy,
  None: noRetryPolicy,
};
>>>>>>> Sort responses
