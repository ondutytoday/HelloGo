package main

import (
	"Module2/collections"
	"Module2/pointers"
	"Module2/structs"
	"fmt"
)

func main() {
	if true {
		collections.ChangeMaps()
	} else {
		pointers.PrintPointers()
		structs.Method()
		var m structs.MyStructure
		m.X = 4
		// m.y = 1 - недоступно т.к с маленькой буквы -> private
		fmt.Println(m)
		collections.Arrays()
		collections.Slices()
		collections.StructSlices()
		collections.ShowMaps()
	}
}
