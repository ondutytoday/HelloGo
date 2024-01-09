package main

import (
	"fmt"
	"strconv"
)

func baseTypes() {
	//bool

	//string

	//int  int8  int16  int32  int64
	//uint uint8 uint16 uint32 uint64 uintptr

	//byte - псевдоним для uint8

	//rune - псевдоним для int32; представляет Unicode код

	//float32 float64

	//complex64 complex128
	var i int
	var f float64
	var b bool
	var c complex64
	s := "string\n"
	fmt.Printf("%v %v %v %v %v\n", i, f, b, c, s)

	var test uint8 = 255
	fmt.Println("Test of max values: " + strconv.Itoa(int(test)))
	test += 2
	fmt.Println("Test of max values: " + strconv.Itoa(int(test)))
	//вывод так же как и в java начинается заново
}
