package main

import (
	"flag"
	"fmt"
	"os"
)

// 自定义命令源码文件的参数使用说明
// 索性不用全局的flag.CommandLine变量，转而自己创建一个私有的命令参数容器

var name string
var cmdLine = flag.NewFlagSet("question(都是我自定义的)", flag.ExitOnError)

func init() {
	cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello, %s!\n", name)
}
