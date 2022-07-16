package main

import (
	"flag"
	"fmt"
	"os"
)

// 自定义命令源码文件的参数使用说明
// 现在再深入一层，我们在调用flag包中的一些函数（比如StringVar、Parse等等）的时候，实际上是在调用flag.CommandLine变量的对应方法。

var name string

func init() {
	flag.CommandLine = flag.NewFlagSet("还是我自定义的", flag.ExitOnError)
	// flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question(这又是我自定义的)")
		flag.PrintDefaults()
	}
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
