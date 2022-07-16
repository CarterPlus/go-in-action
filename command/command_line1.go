package main

import (
	"flag"
	"fmt"
)

// 简单演示命令源码文件接收传入的参数及使用参数

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
