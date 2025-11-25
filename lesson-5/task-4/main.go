package main

import "fmt"

/* ТЗ:
Задание 4. Сделать инверсию мапы.
Например имеем на входе
{
	"a": 1,
	"b": 2,
}
На выходе
{
	1: "a",
	2: "b",
}
*/

// Invert - инвертирует map[string]int в map[int]string
func Invert(m map[string]int) map[int]string {
	resultMap := make(map[int]string)
	for k, v := range m {
		resultMap[v] = k
	}
	return resultMap
}

func main() {
	tests := []map[string]int{
		{"a": 1, "b": 2},
		{"x": 10, "y": 20, "z": 30},
		{"cat": 5, "dog": 7},
	}

	for _, t := range tests {
		fmt.Printf("Исходный map: %v\n", t)
		inv := Invert(t)
		fmt.Printf("Инвертированная: %v\n\n", inv)
	}
}

/* Вывод из консоли:
Исходный map: map[a:1 b:2]
Инвертированная: map[1:a 2:b]

Исходный map: map[x:10 y:20 z:30]
Инвертированная: map[10:x 20:y 30:z]

Исходный map: map[cat:5 dog:7]
Инвертированная: map[5:cat 7:dog]
*/
