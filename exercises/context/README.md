# Context

A `context.Context` carries three things across API boundaries and between goroutines: a cancellation signal, a deadline, and request-scoped values. It is the first parameter of every function that may block or take long (`func Fetch(ctx context.Context, ...)`), and it is passed down, never stored in a struct.

`context.WithCancel`, `WithTimeout` and `WithDeadline` derive a child context from a parent; cancelling the parent cancels all its children. They return a `cancel` function that you should always call (usually with `defer`), or the child lives until its parent ends. Code that waits listens to `ctx.Done()` in a `select`, and `ctx.Err()` says why it ended: `context.Canceled` or `context.DeadlineExceeded`. `WithCancelCause` and `context.Cause` let the canceller give a reason of its own.

`WithValue` attaches a value to a context. Use it only for data that belongs to the request (a request ID, the authenticated user), never for optional parameters, and use a key type of your own so that no other package can collide with it. `context.WithoutCancel` keeps the values but drops the cancellation, and `context.AfterFunc` runs a function when a context ends.

* [Go Concurrency Patterns: Context](https://go.dev/blog/context)
* [Contexts and structs](https://go.dev/blog/context-and-structs)
* [Package context](https://pkg.go.dev/context)
