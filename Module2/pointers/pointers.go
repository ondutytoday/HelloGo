package pointers

import (
	"fmt"
)

// pointer - указатель
// dereferencing - разыменование
// assigning - присваивание
func PrintPointers() {
	i := 42
	p := &i
	fmt.Println(&i)
	fmt.Println(&p)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(&p)
	j := 2790
	*p = 25
	fmt.Println(i)
	p = &j
	fmt.Println(*p)
	*p = *p / 37
	fmt.Println(j)

	fmt.Scan(&p) //нужно записать именно в нашу переменную поэтому &p

	s := "LLL"
	fmt.Println(s)         // LLL
	fmt.Println(change(s)) //LOL
	changeVoid(s)
	fmt.Println(s)        //LLL
	changeWithPointer(&s) //раз метод принимает значение, то должны передать ссылку (т.е. откуда это значение взять)
	fmt.Println(s)        //LOLLOL
	var pointer *string = &s
	changeWithPointer(pointer)
	fmt.Println(s)
}

func change(str string) string {
	str = "LOL"
	return str
}

func changeVoid(str string) {
	//т.к. передаем просто переменную, то значение копируется и сохраняется в новой ячейке памяти
	//т.е. видимо так поддерживается иммутабельность
	str = "LOLLOL"
}

func changeWithPointer(str *string) { //показываем что передаем значение
	*str = "LOLLOL" //меняем именно значение
}
