# Templates

`text/template` fills in a text with data. A template mixes plain text with actions in double braces: `{{.Name}}` writes the field `Name` of the data (the dot), `{{range .Items}} ... {{end}}` repeats its body for each element (with the dot set to the element), `{{if .Admin}} ... {{end}}` makes a choice, and `{{upper .}}` calls a function. Only exported fields can be read, and your own functions have to be registered with `Funcs` before `Parse`. Variables (`{{range $i, $x := .}}`) keep a value for later in the same action.

Parse a template once (`template.Must(template.New("x").Parse(...))` for a template written in the code), and run it with `Execute(writer, data)`, which returns an error when the data doesn't have what the template asks for. For anything that ends up in a web page, use `html/template`: it has the same API, and it escapes the data for the place where it is written.

* [Package text/template](https://pkg.go.dev/text/template)
* [Package html/template](https://pkg.go.dev/html/template)
* [Go by Example: Text Templates](https://gobyexample.com/text-templates)
