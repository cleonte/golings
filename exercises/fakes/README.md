# Fakes and dependency injection

Code that talks to a mail server, a database or the clock is hard to test unless that dependency comes in from outside. *Dependency injection* in Go needs no framework: define a small interface for what the code needs, keep a value of it in a struct field or a parameter, and call that. Production passes the real thing; a test passes a *fake*.

A fake is a simple type you write yourself. It can keep what it was asked to do (a spy), fail on demand (to test the error paths and retries) or work like a tiny in-memory version of the real store. A good fake keeps the contract of the real thing, errors included. Put the interface next to the code that uses it, with only the methods it calls: one-method interfaces are the easiest to fake.

For a global such as `var now = time.Now`, a test can replace it, but must put it back with `t.Cleanup`. Libraries that generate mocks exist, and the standard library has fakes too (`httptest`, `fstest.MapFS`, `iotest`), but hand-written fakes are usually enough.

* [Go Wiki: Test comments](https://go.dev/wiki/TestComments)
* [Go Wiki: Code review comments, interfaces](https://go.dev/wiki/CodeReviewComments#interfaces)
* [Learn Go with Tests: Mocking](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking)
* [Learn Go with Tests: Dependency injection](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection)
