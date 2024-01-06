package loops

import "fmt"

// это чтото типа finally
func TestDefer() {
	fmt.Println("Подсчет...")

	for i := 0; i < 10; i++ {
		//данные записались в стэк и возвращаются из стэка
		defer fmt.Println(i)
	}

	fmt.Println("выполнен")
}
