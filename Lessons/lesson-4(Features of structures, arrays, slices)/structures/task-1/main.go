package main

import "fmt"

/* ТЗ:
Задание 1. Базовая структура и метод со значением

Цель: объявить структуру, понять нулевое значение и метод-значение.
Сделать: `Point{X,Y int}` + метод `Len2() int` (квадрат длины до (0,0)).
Подводные камни: не путайте `int`/`float64`; здесь — целые, квадрат подходит.
Критерии: `Point{3,4}.Len2()==25`, `Point{}.Len2()==0`.
*/

type Point struct {
	X, Y int
}

func (p Point) Len2() int {
	return p.X*p.X + p.Y*p.Y
}

func main() {
	p := Point{3, 4}
	fmt.Println("Point:", p)
	fmt.Println("Квадрат длины:", p.Len2())

	zero := Point{}
	fmt.Println("Нулевая точка:", zero)
	fmt.Println("Квадрат длины (0,0):", zero.Len2())

	// Проверка: (проходит успешно)
	if p.Len2() != 25 {
		panic("want 25")
	}

}

/* Вывод из консоли:
Point: {3 4}
Квадрат длины: 25
Нулевая точка: {0 0}
Квадрат длины (0,0): 0
*/
