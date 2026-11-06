# net/http

A server in Go is an `http.Handler`: a value with `ServeHTTP(w http.ResponseWriter, r *http.Request)`. A `http.ServeMux` picks the handler by pattern; since Go 1.22 a pattern can name a method and contain wildcards (`"GET /items/{id}"`, read with `r.PathValue("id")`). A middleware is a function that wraps a handler and returns a new one.

A response goes out in order: the status, the headers, then the body. The first write to the body sends the status (200 unless you set another one before), so set the status and the headers first. `http.Error` writes an error answer but does not end your function: `return` after it. Limit what you read from a client (`http.MaxBytesReader`) and answer with the right status: 400 for a bad request, 401 for a missing login, 413 for a body that is too big.

`net/http/httptest` tests all of this without a network: `httptest.NewRequest` builds a request, `httptest.NewRecorder` records the answer (`rec.Code`, `rec.Body`). A client, `http.Client`, has no timeout unless you set `Timeout`; give every client you create one.

* [Package net/http](https://pkg.go.dev/net/http)
* [Routing Enhancements for Go 1.22](https://go.dev/blog/routing-enhancements)
* [Package httptest](https://pkg.go.dev/net/http/httptest)
* [The complete guide to Go net/http timeouts](https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/)
