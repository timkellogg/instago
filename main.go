package main

import "github.com/timkellogg/instago/cmd"

var (
	// VERSION set during build
	VERSION = "0.0.1"
)

func main() {
	cmd.Execute(VERSION)
}
