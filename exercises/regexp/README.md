# Regular expressions

The `regexp` package uses the RE2 syntax: linear time, no backtracking, no look-ahead. Compile a pattern once, at package level, with `regexp.MustCompile` (it panics when the pattern is not valid, which is fine for a constant) and use it many times; write patterns as raw strings, between backquotes, so the backslashes stay as they are.

`MatchString` says whether the pattern matches anywhere in the text: add `^` and `$` to require the whole text. `FindString` returns the first match, `FindAllString(s, -1)` all of them, and `FindStringSubmatch` the first match followed by the text of each group in parentheses (nil when there is no match). `ReplaceAllString` replaces matches and can use the groups in the replacement (`${1}`), where `ReplaceAllLiteralString` writes the replacement as it is. Text from outside must go through `regexp.QuoteMeta` before it becomes part of a pattern. For a plain search, `strings.Contains` is simpler and faster.

* [Package regexp](https://pkg.go.dev/regexp)
* [Regular expression syntax](https://pkg.go.dev/regexp/syntax)
* [Go by Example: Regular Expressions](https://gobyexample.com/regular-expressions)
