package main

import "fmt"

/* ТЗ:
Задание 5. Удаление дубликатов из слайса
Описание: Напиши функцию
func unique(nums []int) []int
которая возвращает новый слайс без дубликатов, сохраняя порядок элементов.
Пример:
Вход: [1 2 2 3 1 4 3]
Выход: [1 2 3 4]
Подсказка: Используй map[int]bool для проверки уже встречавшихся элементов, и append() для формирования нового результата.
*/

func main() {
	s := []int{1, 2, 2, 3, 1, 4, 3}
	fmt.Println("Исходный слайс:", s)
	fmt.Println("Слайс без дубликатов:", unique(s))
}

func unique(nums []int) []int {
	seen := make(map[int]bool)
	uniqueSlice := make([]int, 0, len(nums))
	for _, n := range nums {
		if !seen[n] {
			uniqueSlice = append(uniqueSlice, n)
			seen[n] = true
		}
	}
	return uniqueSlice
}

/* Вывод из консоли:
Исходный слайс: [1 2 2 3 1 4 3]
Слайс без дубликатов: [1 2 3 4]
*/
