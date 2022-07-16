package main

import "go-in-action/pkg/d2/lib"

func main() {
	// 在这里调用lib里面的方法，需确保go-in-action/pkg/d2/lib路径下的package也为lib
	lib.Hello("Carter")
}
