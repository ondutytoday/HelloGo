package loops

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func IfOperator(x bool, y, z, lim float64) string {
	if x {
		return "from x"
		//можем инициализировать переменную прямо в блоке и она будет доступна ниже
	} else if v := math.Pow(y, z); v < lim {
		return fmt.Sprintf("v lower than %f", lim)
	} else {
		return fmt.Sprintf("v bigger than %f", lim)
	}
}

func SwitchOperator() string {
	switch os := runtime.GOOS; os {
	case "darwin":
		return "Mac OS"
	case "linux":
		return "Linux"
	default:
		return os
	}
}

func EmptySwitchOperator() string {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		return "Доброе утро!"
	case t.Hour() < 17:
		return "Добрый день."
	default:
		return "Добрый вечер."
	}
}

//должно быть уникальное имя на пакет
/*func ForLoop(x int, y string) int {
	var sum int
	for i := 0; i < x; i++ {
		sum += i
	}
	return sum
}*/
