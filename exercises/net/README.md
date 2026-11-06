# Networking

`net` gives you connections: `net.Dial` opens one, `net.Listen` waits for them, and a `net.Conn` is an `io.Reader` and an `io.Writer`, so `bufio`, `io.Copy` and `fmt.Fprint` all work on it. In tests you don't need the internet: `net.Pipe()` gives two connected ends in memory, and a listener on `127.0.0.1:0` uses a free port on the loopback interface, which every machine has.

A few things about connections are worth knowing early. TCP is a stream, not messages: one `Read` can return part of what the other side wrote, so a protocol needs a delimiter or a length. A reader waits until the other side writes or closes, so use deadlines (`SetDeadline`) and close every connection you accept. Errors are wrapped, so use `errors.Is` and `errors.As`. And a server handles each connection in its own goroutine.

The exercises use only the loopback address and in-memory pipes; they open a port for a moment, which a firewall rarely asks about for `127.0.0.1`.

* [Package net](https://pkg.go.dev/net)
* [Go by Example: TCP servers](https://gobyexample.com/tcp-server)
* [Package net/http/httptest](https://pkg.go.dev/net/http/httptest) (for HTTP, see the `http` topic)
