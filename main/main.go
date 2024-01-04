package main

import (
	"Module1/anotherfile"
	"Module1/fields"
	"Module1/loops"
	"Module1/morestrings"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"rsc.io/quote"
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
	if true {
		loops.ForRangeLoop()
	} else {
		hello := "Hello world!"
		fmt.Println(hello)
		fmt.Println(quote.Go())
		fmt.Println(anotherfile.Version())
		fmt.Println(morestrings.ReverseRunes("some string"))
		fmt.Println(cmp.Diff("Hello World", "Hello Go"))
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
		baseTypes()
		fmt.Println(loops.ForLoop(60))
		fmt.Println(loops.WhileLoop(60))
		fmt.Println(loops.InfiniteLoop(160))
		fmt.Println(loops.IfOperator(true, 1, 1, 1))
		fmt.Println(loops.IfOperator(false, 2, 2, 5))
		fmt.Println(loops.SwitchOperator())
		fmt.Println(loops.EmptySwitchOperator())
		loops.TestDefer()
	}
}
