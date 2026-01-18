package main

/* ТЗ:
Задача 2. Метод с указателем-получателем (изменение состояния)
Условие:
Модифицируй структуру `Point`: добавь метод `Move(dx, dy int)`, который сдвигает точку.
Метод должен изменять сам объект.
*/

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	p := Point{X: 1, Y: 2}
	p.Move(3, -1)
	fmt.Println("Point:", p.X, p.Y) // 4 1
}

/* Вывод из консоли:
Point: 4 1
*/
