package main

import (
	"flag"
	"fmt"
	"os"
)

// 自定义命令源码文件的参数使用说明
// 这有很多种方式，最简单的一种方式就是对变量flag.Usage重新赋值。flag.Usage的类型是func()，即一种无参数声明且无结果声明的函数类型。

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question(这是我自定义的)")
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
