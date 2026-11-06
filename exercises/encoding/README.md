# Encoding

The `encoding` packages turn values into bytes and back, and each has the same two sides: an encoder and a decoder. `encoding/base64` writes bytes as text (use `RawURLEncoding` for anything that goes in a URL or a file name), `encoding/hex` writes them as hexadecimal digits, and both return an error when the text isn't valid, which you must not ignore. `encoding/csv` reads and writes comma-separated values, with quoting and with commas inside quotes handled for you; its writer is buffered, so call `Flush` and check `Error`. `encoding/binary` writes integers as bytes in a fixed order, big-endian or little-endian; two programs that talk to each other must agree on it. JSON has a topic of its own.

* [Package encoding/base64](https://pkg.go.dev/encoding/base64)
* [Package encoding/hex](https://pkg.go.dev/encoding/hex)
* [Package encoding/csv](https://pkg.go.dev/encoding/csv)
* [Package encoding/binary](https://pkg.go.dev/encoding/binary)
