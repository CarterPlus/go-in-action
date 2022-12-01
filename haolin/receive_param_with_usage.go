package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everybody", "The greeting object.")
	flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
}

func main() {
	// flag.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }

	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
