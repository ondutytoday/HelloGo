## Базовые типы Go

- bool
- string
- int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
- byte // псевдоним для uint8
- rune // псевдоним для int32 // представляет Unicode код
- float32 float64
- complex64 complex128

## Нулевое значение:

0 для числовых типов<br>
false для булева типа<br>
"" (пустая строка) для строк

## Форматирование

- %s - строка 
- $d - цифра
- %T - тип переменной
- %v - значение переменной
- %q - строка с эскейпнутыми символами ("string\n" напечатает "string\n", а не "string")

## Неплохо о слайсах
https://habr.com/ru/articles/739754/