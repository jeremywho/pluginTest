package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, _ := plugin.Open("./myplugin.so")
	add, _ := p.Lookup("Add")
	sum := add.(func(int, int) int)(1, 2)
	fmt.Println(sum)
}
