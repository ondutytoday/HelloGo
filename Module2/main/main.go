package main

import (
	"Module2/pointers"
	"Module2/structs"
	"fmt"
)

func main() {
	pointers.PrintPointers()
	structs.Method()
	var m structs.MyStructure
	m.X = 4
	// m.y = 1 - недоступно т.к с маленькой буквы -> private
	fmt.Println(m)
}
