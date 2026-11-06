package order

var config string

// TODO: a_server.go comes before b_config.go, so its init runs first and sees an empty config. Give config its value in its declaration instead: all variables of a package are initialized before the first init function runs.
func init() {
	config = ":8080"
}
