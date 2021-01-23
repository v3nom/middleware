# middleware [![Build Status](https://travis-ci.com/v3nom/middleware.svg?branch=master)](https://travis-ci.com/v3nom/middleware)
Useful middlewares for [Pipes](https://github.com/v3nom/pipes) library.

## What's in the box?

### Context middleware
Creates request Context and passes it forward.

### Panic recovery middleware
Handles panic and allows to provide user friendly message or redirect to error page using http.ResponseWriter.

### Rate limiter middleware (NOT TESTED IN PRODUCTION)
Limits network requests from single IP using provided time function. 

### Cookie auth middleware
Decodes user object from requests cookie and adds to context
