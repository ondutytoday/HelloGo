package loops

import (
	"fmt"
)

func ForLoop(x int) int {
	var sum int
	for i := 0; i < x; i++ {
		sum += i
	}
	return sum
}

func WhileLoop(x int) int {
	sum := 1
	//it is like for(;sum<x;){}
	for sum < x {
		sum += sum
	}
	return sum
}

func InfiniteLoop(x int) int {
	sum := 1
	for {
		sum += sum
		if sum > x {
			break
		}
	}
	return sum
}

func ForRangeLoop() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	//проходим по листу
	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}
	//если не нужен итератор
	for _, v := range pow {
		fmt.Printf("2 в квадрате это %d\n", v)
	}
}
