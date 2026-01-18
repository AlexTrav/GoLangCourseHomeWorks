package main

import "fmt"

/* ТЗ:
Задание 2. Конструктор и мутация по указателю

Цель: создать конструктор, мутацию состояния, понять когда нужен указатель.
Сделать: `NewPoint(x,y int) *Point`; метод `Move(dx,dy int)` на `*Point`.
Теория: указатель нужен для изменения состояния и чтобы не копировать большие структуры.
Критерии: старт `(1,1)`, `Move(2,-1)` → `(3,0)`.
*/

type Point struct {
	X, Y int
}

func NewPoint(x, y int) *Point {
	return &Point{X: x, Y: y}
}

func (p *Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	p := NewPoint(1, 1)
	fmt.Println("Начальная точка:", *p)

	p.Move(2, -1)
	fmt.Println("После Move(2, -1):", *p)

	// Проверка: (проходит успешно)
	if p.X != 3 || p.Y != 0 {
		panic("bad move")
	}
}

/* Вывод из консоли:
Начальная точка: {1 1}
После Move(2, -1): {3 0}
*/
