<a name="module_swagger-test"></a>

## swagger-test
swagger-test client library.


* [swagger-test](#module_swagger-test)
    * [SwaggerTest](#exp_module_swagger-test--SwaggerTest) ⏏
        * [new SwaggerTest(options)](#new_module_swagger-test--SwaggerTest_new)
        * _instance_
            * [.getBook(id, [options], [cb])](#module_swagger-test--SwaggerTest+getBook) ⇒ <code>Promise</code>
        * _static_
            * [.RetryPolicies](#module_swagger-test--SwaggerTest.RetryPolicies)
                * [.Exponential](#module_swagger-test--SwaggerTest.RetryPolicies.Exponential)
                * [.Single](#module_swagger-test--SwaggerTest.RetryPolicies.Single)
                * [.None](#module_swagger-test--SwaggerTest.RetryPolicies.None)
            * [.Errors](#module_swagger-test--SwaggerTest.Errors)
                * [.ExtendedError](#module_swagger-test--SwaggerTest.Errors.ExtendedError) ⇐ <code>Error</code>
                * [.NotFound](#module_swagger-test--SwaggerTest.Errors.NotFound) ⇐ <code>Error</code>
                * [.InternalError](#module_swagger-test--SwaggerTest.Errors.InternalError) ⇐ <code>Error</code>

<a name="exp_module_swagger-test--SwaggerTest"></a>

### SwaggerTest ⏏
swagger-test client

**Kind**: Exported class  
<a name="new_module_swagger-test--SwaggerTest_new"></a>

#### new SwaggerTest(options)
Create a new client object.


| Param | Type | Default | Description |
| --- | --- | --- | --- |
| options | <code>Object</code> |  | Options for constructing a client object. |
| [options.address] | <code>string</code> |  | URL where the server is located. Must provide this or the discovery argument |
| [options.discovery] | <code>bool</code> |  | Use @clever/discovery to locate the server. Must provide this or the address argument |
| [options.timeout] | <code>number</code> |  | The timeout to use for all client requests, in milliseconds. This can be overridden on a per-request basis. |
| [options.retryPolicy] | <code>[RetryPolicies](#module_swagger-test--SwaggerTest.RetryPolicies)</code> | <code>RetryPolicies.Single</code> | The logic to determine which requests to retry, as well as how many times to retry. |

<a name="module_swagger-test--SwaggerTest+getBook"></a>

#### swaggerTest.getBook(id, [options], [cb]) ⇒ <code>Promise</code>
**Kind**: instance method of <code>[SwaggerTest](#exp_module_swagger-test--SwaggerTest)</code>  
**Fulfill**: <code>undefined</code>  
**Reject**: <code>[ExtendedError](#module_swagger-test--SwaggerTest.Errors.ExtendedError)</code>  
**Reject**: <code>[NotFound](#module_swagger-test--SwaggerTest.Errors.NotFound)</code>  
**Reject**: <code>[InternalError](#module_swagger-test--SwaggerTest.Errors.InternalError)</code>  
**Reject**: <code>Error</code>  

| Param | Type | Description |
| --- | --- | --- |
| id | <code>number</code> |  |
| [options] | <code>object</code> |  |
| [options.timeout] | <code>number</code> | A request specific timeout |
| [options.span] | <code>[Span](https://doc.esdoc.org/github.com/opentracing/opentracing-javascript/class/src/span.js~Span.html)</code> | An OpenTracing span - For example from the parent request |
| [options.retryPolicy] | <code>[RetryPolicies](#module_swagger-test--SwaggerTest.RetryPolicies)</code> | A request specific retryPolicy |
| [cb] | <code>function</code> |  |

<a name="module_swagger-test--SwaggerTest.RetryPolicies"></a>

#### SwaggerTest.RetryPolicies
Retry policies available to use.

**Kind**: static property of <code>[SwaggerTest](#exp_module_swagger-test--SwaggerTest)</code>  

* [.RetryPolicies](#module_swagger-test--SwaggerTest.RetryPolicies)
    * [.Exponential](#module_swagger-test--SwaggerTest.RetryPolicies.Exponential)
    * [.Single](#module_swagger-test--SwaggerTest.RetryPolicies.Single)
    * [.None](#module_swagger-test--SwaggerTest.RetryPolicies.None)

<a name="module_swagger-test--SwaggerTest.RetryPolicies.Exponential"></a>

##### RetryPolicies.Exponential
The exponential retry policy will retry five times with an exponential backoff.

**Kind**: static constant of <code>[RetryPolicies](#module_swagger-test--SwaggerTest.RetryPolicies)</code>  
<a name="module_swagger-test--SwaggerTest.RetryPolicies.Single"></a>

##### RetryPolicies.Single
Use this retry policy to retry a request once.

**Kind**: static constant of <code>[RetryPolicies](#module_swagger-test--SwaggerTest.RetryPolicies)</code>  
<a name="module_swagger-test--SwaggerTest.RetryPolicies.None"></a>

##### RetryPolicies.None
Use this retry policy to turn off retries.

**Kind**: static constant of <code>[RetryPolicies](#module_swagger-test--SwaggerTest.RetryPolicies)</code>  
<a name="module_swagger-test--SwaggerTest.Errors"></a>

#### SwaggerTest.Errors
Errors returned by methods.

**Kind**: static property of <code>[SwaggerTest](#exp_module_swagger-test--SwaggerTest)</code>  

* [.Errors](#module_swagger-test--SwaggerTest.Errors)
    * [.ExtendedError](#module_swagger-test--SwaggerTest.Errors.ExtendedError) ⇐ <code>Error</code>
    * [.NotFound](#module_swagger-test--SwaggerTest.Errors.NotFound) ⇐ <code>Error</code>
    * [.InternalError](#module_swagger-test--SwaggerTest.Errors.InternalError) ⇐ <code>Error</code>

<a name="module_swagger-test--SwaggerTest.Errors.ExtendedError"></a>

##### Errors.ExtendedError ⇐ <code>Error</code>
ExtendedError

**Kind**: static class of <code>[Errors](#module_swagger-test--SwaggerTest.Errors)</code>  
**Extends:** <code>Error</code>  
**Properties**

| Name | Type |
| --- | --- |
| code | <code>number</code> | 
| message | <code>string</code> | 

<a name="module_swagger-test--SwaggerTest.Errors.NotFound"></a>

##### Errors.NotFound ⇐ <code>Error</code>
NotFound

**Kind**: static class of <code>[Errors](#module_swagger-test--SwaggerTest.Errors)</code>  
**Extends:** <code>Error</code>  
**Properties**

| Name | Type |
| --- | --- |
| message | <code>string</code> | 

<a name="module_swagger-test--SwaggerTest.Errors.InternalError"></a>

##### Errors.InternalError ⇐ <code>Error</code>
InternalError

**Kind**: static class of <code>[Errors](#module_swagger-test--SwaggerTest.Errors)</code>  
**Extends:** <code>Error</code>  
**Properties**

| Name | Type |
| --- | --- |
| code | <code>number</code> | 
| message | <code>string</code> | 
