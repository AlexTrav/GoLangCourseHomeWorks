package main

import "fmt"

/* ТЗ:
Задание 2. Умножение всех элементов слайса
Описание:Создай слайс из 5 чисел. Реализуй функцию multiply(s []int, factor int) []int,
которая умножает каждый элемент на заданный множитель.
Пример:
Вход: [2 4 6], factor = 3
Выход: [6 12 18]
Подсказка: Слайс изменяется по ссылке, если его не копировать.
Попробуй обе версии: с изменением на месте и с копией (copy()).
*/

func main() {
	s1 := []int{2, 4, 6, 8, 10}
	fmt.Printf("multiply(s1): %v\n", multiply(s1, 3))

	s2 := []int{2, 4, 6, 8, 10}
	fmt.Printf("multiplyCopy(s2): %v\n", multiplyCopy(s2, 3))

	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v", s2)
}

func multiply(slice []int, factor int) []int {
	resultSlice := slice
	for i := 0; i < len(slice); i++ {
		resultSlice[i] *= factor
	}
	return resultSlice
}

func multiplyCopy(slice []int, factor int) []int {
	resultSlice := make([]int, len(slice))
	copy(resultSlice, slice)
	for i := 0; i < len(slice); i++ {
		resultSlice[i] *= factor
	}
	return resultSlice
}

/* Вывод из консоли:
multiply(s1): [6 12 18 24 30]
multiplyCopy(s2): [6 12 18 24 30]
s1: [6 12 18 24 30]
s2: [2 4 6 8 10]
*/
