package main

import (
	"HelloGo/anotherfile"
	"HelloGo/morestrings"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"rsc.io/quote"
)

func main() {
	hello := "Hello world!"
	fmt.Println(hello)
	fmt.Println(quote.Go())
	fmt.Println(anotherfile.Version())
	fmt.Println(morestrings.ReverseRunes("some string"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
