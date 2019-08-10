# middleware [![Build Status](https://travis-ci.com/v3nom/middleware.svg?branch=master)](https://travis-ci.com/v3nom/middleware)
Useful middlewares for [Pipes](https://github.com/v3nom/pipes) library. Most middlewares are made specifically for the Google App Engine environment. Others are universally useful.

## What's in the box?

### Context middleware
Creates Google App Engine Context and passes it forward.

### Panic recovery middleware
Handles panic and allows to provide user friendly message or redirect to error page using http.ResponseWriter.

### Rate limiter middleware (NOT TESTED IN PRODUCTION)
Limits network requests from single IP using provided time function. 
