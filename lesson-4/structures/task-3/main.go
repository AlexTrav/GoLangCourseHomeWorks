package main

import "fmt"

/* ТЗ:
Задание 3. Композиция: прямоугольник из двух точек

Цель: использовать другие структуры как поля и вычислить свойства.
Сделать: `Rect{Min, Max Point}`; методы `Width()`, `Height()`, `Area()`.
Ограничения: ширина/высота = `abs(Max - Min)` по осям, площадь ≥ 0.
Подводные камни: отрицательные размеры если Min>Max — исправляйте `abs`.
Критерии: `Rect{(0,0),(3,2)}.Area()==6`.
*/

type Point struct {
	X, Y int
}

type Rect struct {
	Min, Max Point
}

func (r Rect) Width() int {
	d := r.Max.X - r.Min.X
	if d < 0 {
		d = -d
	}
	return d
}

func (r Rect) Height() int {
	d := r.Max.Y - r.Min.Y
	if d < 0 {
		d = -d
	}
	return d
}

func (r Rect) Area() int {
	return r.Width() * r.Height()
}

func main() {
	r := Rect{
		Min: Point{0, 0},
		Max: Point{3, 2},
	}

	fmt.Printf("Прямоугольник: %+v\n", r)
	fmt.Println("Ширина:", r.Width())
	fmt.Println("Высота:", r.Height())
	fmt.Println("Площадь:", r.Area())

	// Проверка: (проходит успешно)
	if r.Area() != 6 {
		panic("Ошибка: ожидалась площадь 6")
	}
}

/* Вывод из консоли:
Прямоугольник: {Min:{X:0 Y:0} Max:{X:3 Y:2}}
Ширина: 3
Высота: 2
Площадь: 6
*/
