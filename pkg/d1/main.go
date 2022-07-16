package main

import (
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	hello(name) // go run pkg/d1/main.go pkg/d1/lib.go --name Carter
}
