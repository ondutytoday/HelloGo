package main

import (
	"fmt"
	"strconv"
)

func swap(x, y string) (string, string) {
	return y, x
}

func namedReturn(x, y int) (a int, b string) {
	a = x
	b = "string named: " + strconv.Itoa(y) //converts int value to string
	return
}

func main() {
	//hello := "Hello world!"
	//fmt.Println(hello)
	//fmt.Println(quote.Go())
	//fmt.Println(anotherfile.Version())
	//fmt.Println(morestrings.ReverseRunes("some string"))
	//fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println(swap("a", "b"))
	a, b := namedReturn(3, 5)
	fmt.Printf("return a: %d\nreturn b: %s\n", a, b)
	s := string(1024) //converts bytes to char
	fmt.Println(s)
}
