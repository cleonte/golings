# io and bufio

Most of Go's input and output is built from two small interfaces: `io.Reader` (`Read(p []byte) (n int, err error)`) and `io.Writer` (`Write(p []byte) (n int, err error)`). Files, network connections, buffers, compressors and hashes implement them, so code written against the interfaces works with all of them, and tests can use a `strings.NewReader` or a `bytes.Buffer` instead of a real file.

A `Read` fills as much of `p` as it can and returns how many bytes it put there, which may be fewer than `len(p)`. At the end of the input it returns `io.EOF`. Use `io.ReadAll` to read everything, `io.Copy` to move data from a reader to a writer, and `io.LimitReader`, `io.TeeReader` and `io.MultiWriter` to combine and restrict them. `fmt.Fprintf(w, ...)` formats to any writer.

`bufio` adds buffering. `bufio.Scanner` reads lines (or words) and stops at a line over 64 KiB, so check `scanner.Err()` after the loop; `bufio.NewReader` and `bufio.NewWriter` reduce the number of system calls (call `Flush` on a writer).

* [Package io](https://pkg.go.dev/io)
* [Package bufio](https://pkg.go.dev/bufio)
* [Go by Example: Reading Files](https://gobyexample.com/reading-files)
* [Go by Example: Line Filters](https://gobyexample.com/line-filters)
