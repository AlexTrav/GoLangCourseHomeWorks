package main

import "fmt"

/* ТЗ:
Задание 3. Обрезка и копирование
Описание:Создай слайс numbers := make([]int, 0, 20) и добавь в него числа 1–10.
Затем создай под-слайс part := numbers[3:8].
Создай копию copyPart, чтобы освободиться от исходного массива.
Измени элементы в part и убедись, что это не влияет на copyPart.
Выведи:
длину и ёмкость part и copyPart
адрес первого элемента (через &part[0]) для сравнения
Цель: Понять, когда слайсы указывают на одну память, а когда нет.
*/

func main() {
	numbers := make([]int, 0, 20)
	numbers = append(numbers, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	part := numbers[3:8]

	copyPart := make([]int, len(part))
	copy(copyPart, part)

	part[0] = 11
	part[1] = 12
	part[2] = 13
	part[3] = 14

	fmt.Printf("part: %v; len: %v; cap: %v\n", part, len(part), cap(part))
	fmt.Printf("&part[0]: %p\n", &part[0])

	fmt.Printf("copyPart: %v; len: %v; cap: %v\n", copyPart, len(copyPart), cap(copyPart))
	fmt.Printf("&copyPart[0]: %p\n", &copyPart[0])
}

/* Вывод из консоли:
part: [11 12 13 14 8]; len: 5; cap: 17
&part[0]: 0xc000122018
copyPart: [4 5 6 7 8]; len: 5; cap: 5
&copyPart[0]: 0xc000124000
*/
