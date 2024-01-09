package collections

import "fmt"

type MyStruct struct {
	Name, Surname string
}

func ChangeMaps() {
	m := map[string]string{
		"one":  "two",
		"tree": "four",
	}
	//put
	m["five"] = "six"
	//get
	a := m["one"]
	//remove
	delete(m, "three")
	//if the key exists
	b, ok := m["dsvsdf"]

	fmt.Println(m, a, b, ok)
}

func ShowMaps() {
	var myMap map[string]string
	myMap = make(map[string]string)
	myMap["one"] = "two"

	fmt.Println(myMap["one"], myMap["two"], myMap)

	myMap2 := make(map[string]MyStruct)
	myMap3 := make(map[MyStruct]string)

	myStruct1 := MyStruct{"Myone", "Mytwo"}
	myStruct2 := MyStruct{"Myfour", "Myfive"}
	myMap2["one"] = myStruct1
	myMap2["two"] = myStruct2

	myMap3[myStruct1] = "three"
	myMap3[myStruct2] = "six"
	fmt.Println(myMap2, myMap3)

	myMap3[myStruct1] = "seven"
	myMap3[MyStruct{"Myfour", "Myfive"}] = "nine"
	fmt.Println(myMap2, myMap3)

	//поменяли значение в ключе -> значение найдено не будет
	myStruct2.Name = "changed"
	fmt.Printf("print: %v\n", myStruct2)
	fmt.Println("print: " + myMap3[myStruct2])
	fmt.Println(myMap3)

	//еще один вид инициализации
	var m = map[string]MyStruct{
		"Bell Labs": {"40.68433", "-74.39967"},
		"Google":    {"37.42202", "-122.08408"},
	}
	fmt.Println(m)
}
