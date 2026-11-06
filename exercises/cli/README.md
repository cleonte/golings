# Command-line programs

The `flag` package reads the command line: declare flags on a `flag.FlagSet` (`fs.String`, `fs.Int`, `fs.Bool`, `fs.Func`...), call `fs.Parse(args)`, and read the values; `fs.Args()` holds what is left after the flags. A program normally uses the default set through `flag.Parse()` and `os.Args`, but a function that takes `args []string` as a parameter is just as easy to write and much easier to test.

The same idea applies to the whole program: write `run(args []string, stdout, stderr io.Writer) int` that returns an exit code, and make `main` one line, `os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))`. Results go to stdout, problems to stderr; exit code 0 is success, 1 a failure and 2 a wrong command line. Tools with commands give each command its own flag set.

* [Package flag](https://pkg.go.dev/flag)
* [Package os: Args and Exit](https://pkg.go.dev/os#Exit)
* [Go by Example: Command-Line Flags](https://gobyexample.com/command-line-flags)
* [Go by Example: Command-Line Subcommands](https://gobyexample.com/command-line-subcommands)
