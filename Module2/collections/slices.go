package collections

import "fmt"

//-----------------------------------------------------------
//		СНОСКА 1:
//вместимость говорит о том сколько памяти нужно выделить заранее под наш массив,
//чтобы при расширении нам не пришлось искать новый участок памяти
//если не указать, то будет == длине
//если превысить то вместимость*2 как в ArrayList
//		При этом в памяти произойдет следующее:
//1. Go понимает что нам не хватает памяти и посмотрит есть ли после текущего сегмента памяти еще столько же ячеек;
//2. Если ячейки есть, он не будет передвигать массив и просто зарезервирует больше памяти;
//3. Если ячеек нет, то он скопирует всю информацию из уже использующегося сегмента
//и найдет вдвое больше свободных ячеек, после нахождения он перенесет туда все данные
//и отдаст нам адрес сегмента;

//-----------------------------------------------------------
//		СНОСКА 2:
//  Срез не хранит никаких данных, он всего лишь обозначает секцию нижележащего массива.(надмножество массива)
//	Изменения элементов среза приводят к модификации соответствующих элементов его нижележащего массива.
//	Другие срезы, имеющие общий нижележащий массив, также увидят эти изменения.

func Slices() {
	var a []int
	//a[0] = 15 - так нельзя. Получится IndexOutOfBounds
	//СНОСКА 1
	a = make([]int, 3, 5) //тип, длина, вместимость(необязательный параметр)
	fmt.Println(a)
	//при использовании функции append - мы переприсваиваем значение переменной
	//Если нижележащий массив среза s слишком мал, чтобы вместить все значения,
	//то будет создан новый массив большего размера. Результирующий срез будет ссылаться на этот новый массив.
	a = append(a, 1)
	fmt.Println(a)
	a = append(a, 3)
	fmt.Println(a)
	a = append(a, 6) //вышли за пределы обозначенной вместимости
	fmt.Println(a)

	aa := [2]int{1, 2}
	aaa := aa[:]
	fmt.Println("Test aa, aaa", aa, aaa)
	fmt.Println("Test &aa, &aaa", &aa, &aaa)
	aaa[0] = 15
	fmt.Println("Test Change: Test aa, aaa", aa, aaa)
	fmt.Println("Test Change: Test &aa, &aaa", &aa, &aaa)
	bb := &aa
	bbb := &aaa
	fmt.Println("Test bb, bbb", &bb, *bb, &bbb, *bbb)
	//НЕ САМАЯ ОЧЕВИДНАЯ И ПОНЯТНАЯ ХРЕНЬ:
	//тут мы превысили вместимость нашего прежнего массива а значит создали новый массив, а значит
	// старый у нас останется дальше неизмененным, а новый уже будем менять
	aaa = append(aaa, 17)
	fmt.Println("After: Test aa, aaa", aa, aaa)
	fmt.Println("After: Test &aa, &aaa", &aa, &aaa)
	fmt.Println("After: Test bb, bbb", &bb, *bb, &bbb, *bbb)
	aaa[0] = 11
	fmt.Println("After1: Test aa, aaa", aa, aaa)
	fmt.Println("After1: Test &aa, &aaa", &aa, &aaa)
	fmt.Println("After1: Test bb, bbb", &bb, *bb, &bbb, *bbb)

	//нулевые значения
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}

	// 0, 1, 2, 3, 4
	b := []int{1, 2, 3, 4, 5}
	//b[10] = 15 - IndexOutOfBounds
	b = append(b, 15)

	c := b[1:4]          //с 1 включительно по 4 невключительно [1,2,3,4) -> 2,3,4
	var d []int = c[0:2] // -> 2, 3
	//a[0:10]
	//a[:10] = [0:10]
	//a[0:] = [0:end]
	//a[:] = [begin:end]
	firstThreeInts := b[:3] // = [0:3]
	fmt.Println(firstThreeInts)
	fmt.Println(b, c, d) // [1 2 3 4 5 15] [2 3 4] [2 3]

	//СНОСКА 2
	d[0] = 100
	fmt.Println(b, c, d) //[1 100 3 4 5 15] [100 3 4] [100 3]

	//array to slice
	nameArray := [6]string{"D", "a", "n", "i", "i", "l"}
	nameSlice := nameArray[:]
	nameSlice = append(nameSlice, "!")
	fmt.Println(nameSlice)

	arraysAndSlices()
	copySlice()
}

func arraysAndSlices() {
	//создаем массив
	nameArray := [6]string{"D", "a", "n", "i", "i", "l"}
	//на его основе создаем срез (slice)
	nameSlice := nameArray[:3]
	//меняем одну позицию в срезе (а следовательно и в массиве)
	nameSlice[len(nameSlice)-1] = "m"
	fmt.Println(nameSlice) // [D a m]
	//хотя срез имеет длину 3, но его вместимость 6 - по длине изначального массива
	fmt.Println(len(nameSlice), cap(nameSlice))
	// Делаем новый срез из старого, но с длиной == вместимости и получаем наш начальный массив
	nameSlice = nameSlice[0:cap(nameSlice)]
	fmt.Println(nameSlice) // [D a m i i l]
}

// если мы просто напишем secondNameSlice := nameSlice,
// то тогда присвоится ссылка и фактически это будет один массив
func copySlice() {
	nameSlice := []string{"D", "a", "n", "i", "i", "l"}
	secondNameSlice := make([]string, len(nameSlice), cap(nameSlice))
	//В случае с copy мы передаем саму переменную (не ссылку, а именно переменную!)
	copy(secondNameSlice, nameSlice)
	secondNameSlice[0] = "T"

	fmt.Println(nameSlice, secondNameSlice) // [D a n i i l] [T a n i i l]
}
