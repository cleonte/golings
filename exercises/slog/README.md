# Structured logging with log/slog

`log/slog` (Go 1.21) writes log records made of a level, a message and key/value pairs: `logger.Info("user logged in", "user", "ana", "attempts", 3)`. Pairs, not a sentence built with `Sprintf`, can be searched and filtered by the tools that read the logs. A handler decides the format (`slog.NewTextHandler` writes `key=value`, `slog.NewJSONHandler` writes JSON) and the lowest level kept, set with `HandlerOptions.Level` (the default is Info). `logger.With(...)` returns a logger that adds pairs to every record, `slog.Group` nests pairs under a name, and `HandlerOptions.ReplaceAttr` rewrites or drops attributes before they are written, to hide a secret or leave out the time.

* [Package log/slog](https://pkg.go.dev/log/slog)
* [Structured Logging with slog](https://go.dev/blog/slog)
* [A Guide to Writing slog Handlers](https://go.dev/s/slog-handler-guide)
