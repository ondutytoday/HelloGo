package structs

import "fmt"

type MyStructure struct {
	X int
	y int
}

func Method() {
	m := MyStructure{1, 2}
	m.X = 5
	m.y = 6
	fmt.Println(m)
	p := &m
	fmt.Println(p) //для структур отображается в виде &{5 6} а не ссылка на ячейку
	fmt.Println(*p)
	p.X = 78
	fmt.Println(p)
	fmt.Println(m)

	var ( //можем определить сразу несколько переменных, чтобы красивее и читабельнее
		v1 = MyStructure{1, 2}  // имеет type MyStructure
		v2 = MyStructure{X: 1}  // не указали одну из переменных -> она 0 или еще какое-то значение по умолчанию
		v3 = MyStructure{}      //X:0 и Y:0
		p1 = &MyStructure{3, 4} // имеет type *MyStructure
		// почему-то считается что тип идет со *, пока не поняла смысл. По факту p1 - это ссылка на объект, а не сам объект
	)

	fmt.Println(v1, p1, v2, v3)
}
