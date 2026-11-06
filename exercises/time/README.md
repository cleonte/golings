# Time

`time.Time` is an instant, with a location for displaying it; `time.Duration` is a length of time, a count of nanoseconds with its own type (`3 * time.Second`, and `time.Duration(n) * time.Second` when `n` is a variable). `t.Add(d)` moves a time by a duration, `t.Sub(u)` gives the duration between two times, and `t.AddDate(y, m, d)` moves by calendar units.

Go formats and parses times with an example instead of codes: the layout is the reference time `Mon Jan 2 15:04:05 MST 2006`, written the way you want the result to look (`2006-01-02` for a date, `15:04` for a 24-hour clock). `time.Parse` reads a time without a zone as UTC; `time.ParseInLocation` reads it in a zone you choose.

Compare instants with `Equal`, `Before` and `After`, not `==`, which also compares the location. `Truncate` and `Round` count from the zero time in UTC, so they are not the way to get "midnight here": build it with `time.Date`. A `time.Timer` fires once and a `time.Ticker` repeatedly; `AfterFunc` runs a function when its timer fires, and `Reset` moves the timer. Tests that wait for time should run in `testing/synctest`, where the clock is fake.

* [Package time](https://pkg.go.dev/time)
* [Go by Example: Time Formatting / Parsing](https://gobyexample.com/time-formatting-parsing)
* [Go by Example: Epoch](https://gobyexample.com/epoch)
