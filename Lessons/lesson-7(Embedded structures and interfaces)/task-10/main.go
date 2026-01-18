package main

/* ТЗ:
Задача 10. Пустой интерфейс (any) и type switch
Условие:
Напиши функцию `Describe(v any)`, которая:
- печатает тип (`int`, `string`, `bool`, `unknown`) и значение,
- использует `type switch`.
Протестируй на нескольких типах.
*/

import "fmt"

func Describe(v any) {
	switch val := v.(type) {
	case int:
		fmt.Println("int:", val)
	case string:
		fmt.Println("string:", val)
	case bool:
		fmt.Println("bool:", val)
	default:
		fmt.Printf("unknown (%T): %v\n", val, val)
	}
}

func main() {
	Describe(10)          // int
	Describe("hello")     // string
	Describe(true)        // bool
	Describe(3.14)        // float64 (unknown)
	Describe([]int{1, 2}) // slice (unknown)
	Describe(nil)         // <nil> (unknown)
}

/* Вывод из консоли:
int: 10
string: hello
bool: true
unknown (float64): 3.14
unknown ([]int): [1 2]
unknown (<nil>): <nil>
*/
