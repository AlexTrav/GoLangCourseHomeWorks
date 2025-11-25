package main

import "fmt"

/* ТЗ:
Задание 5. Слияние 2-х мап
На вход имеем {"a":3, "c":5, "f": 3}, {"f":2, "d": 4, "a": 5}
На выход получаем {"a":8, "c": 5, "f": 5, "d": 4}
*/

// MergeMaps - сливает две map[string]int, суммируя совпадающие ключи
func MergeMaps(dst, src map[string]int) map[string]int {
	resultMap := make(map[string]int)
	for key, value := range dst {
		resultMap[key] = value
	}
	for key, value := range src {
		resultMap[key] += value
	}
	return resultMap
}

func main() {
	tests := [][2]map[string]int{
		{
			{"a": 3, "c": 5, "f": 3},
			{"f": 2, "d": 4, "a": 5},
		},
		{
			{"x": 10, "y": 20},
			{"y": 5, "z": 7},
		},
		{
			{}, {"a": 1, "b": 2},
		},
		{
			{"a": 1, "b": 2}, {"a": 1, "b": 2},
		},
	}

	for _, t := range tests {
		dst := t[0]
		src := t[1]

		fmt.Printf("map1: %v\nmap2: %v\n", dst, src)
		merged := MergeMaps(dst, src)
		fmt.Printf("Результат: %v\n\n", merged)
	}
}

/* Вывод из консоли:
map1: map[a:3 c:5 f:3]
map2: map[a:5 d:4 f:2]
Результат: map[a:8 c:5 d:4 f:5]

map1: map[x:10 y:20]
map2: map[y:5 z:7]
Результат: map[x:10 y:25 z:7]

map1: map[]
map2: map[a:1 b:2]
Результат: map[a:1 b:2]

map1: map[a:1 b:2]
map2: map[a:1 b:2]
Результат: map[a:2 b:4]
*/
