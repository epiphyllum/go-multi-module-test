package main

import (
	"github.com/epiphyllum/go-multi-module-test/hello/algo"
	"github.com/epiphyllum/go-multi-module-test/hello/funcs"
)

func main() {
	add := algo.Add(10, 10)
	println(add)

	funcs.F2()
	funcs.F1()

}
