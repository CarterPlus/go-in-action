package main

import "fmt"

// 学习变量
func test() {
	var name string

	var (
		a int
		b string
		c []float32
		d func() bool
		e struct{}
	)

	fmt.Printf(name, "testing", a, b, c, d, e)
}
