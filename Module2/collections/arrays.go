package collections

import "fmt"

// массив в Go является примитивным типом данных,
// он может копироваться при передаче в другую переменную.
// По умолчанию в Go все значения копируются, а не передаются с помощью ссылки.
// Это значит, что если мы передадим наш массив в функцию,
// то Go скопирует данный массив и в функции будет находиться уже совершенно другой массив
// (вернее точная копия исходного массива)
func Arrays() {
	var a [5]string
	a[1] = "one"
	b := [3]int{1, 2, 3}

	c := [3]int{}
	d := [2]int{}
	// (c) [2]int и (d) [3]int - разные типы, так как длина массива в Go, входит в его тип
	e := [...]int{1, 2, 3} // [3]int

	fmt.Println(a, b, c, d, e)

	//значение и ссылка
	var initArray = [...]int{1, 2, 3}
	var copyArray = initArray
	fmt.Printf("Address of initArray: %p\n", &initArray)
	fmt.Printf("Address of copyArray: %p\n", &copyArray)
}
