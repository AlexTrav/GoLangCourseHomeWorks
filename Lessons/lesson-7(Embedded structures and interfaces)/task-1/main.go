package main

/* ТЗ:
Задача 1. Простейший метод со значением-получателем
Условие:
Создай структуру `Point` с полями `X` и `Y` (int).
Добавь метод `DistanceFromOrigin() int`, который возвращает расстояние до начала координат по формуле `|X| + |Y|` (манхэттен).
В `main` создай точку и выведи расстояние.
*/

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func (p Point) DistanceFromOrigin() int {
	return int(math.Abs(float64(p.X)) + math.Abs(float64(p.Y)))
}

func main() {
	p := Point{X: -3, Y: 5}
	fmt.Println("Distance:", p.DistanceFromOrigin())
}

/* Вывод из консоли:
Distance: 8
*/
