package main

import (
	"HelloGo/fields"
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
	fmt.Println(fields.TestFields())
	//приведение типов
	f := 42.6
	i := int(f) //42
	//ss := string(f) - error
	fmt.Println(i)
	v := 324.4324 + 32i //complex
	fmt.Printf("v is of type %T\n", v)
	fmt.Println(fields.Pi)
	fmt.Println(fields.NeedInt(fields.Small))
	fmt.Println(fields.NeedFloat(fields.Small))
	fmt.Println(fields.NeedFloat(fields.Big))
}
