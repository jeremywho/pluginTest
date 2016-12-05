package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("./myplugin.so")
	failOnError(err)
	add, err := p.Lookup("Add")
	failOnError(err)
	sum := add.(func(int, int) int)(1, 2)
	fmt.Println(sum)
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}
