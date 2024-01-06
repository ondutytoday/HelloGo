package fields

import (
	"strconv"
)

var x, y int

const Pi = 3.14
const (
	// Создаем большое число,
	// сдвигая 1 бит влево на 100 позиций.
	// Другими словами, бинарное число
	// состоящее из 1 и 100 нулей следом.
	Big = 1 << 100
	// Сдвигаем вправо снова 99 позиций,
	// таким образом мы получим 1<<1, или 2.
	Small = Big >> 99
)

func NeedInt(x int) int { return x*10 + 1 }
func NeedFloat(x float64) float64 {
	return x * 0.1
}

func TestFields() string {
	x = 34
	return anotherFunc()
}

func anotherFunc() string {
	x = 14
	return strconv.Itoa(x)
}
