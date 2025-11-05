package main

// ТЗ - https://go.dev/tour/flowcontrol/8

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0 // начальное предположение
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z) // формула Ньютона
	}
	return z
}

func main() {
	fmt.Println("My Sqrt(16) =", Sqrt(16))
	fmt.Println("math.Sqrt(16) =", math.Sqrt(16))

	fmt.Println("My Sqrt(2)  =", Sqrt(2))
	fmt.Println("math.Sqrt(2) =", math.Sqrt(2))
}

/* Вывод из консоли:
My Sqrt(16) = 4
math.Sqrt(16) = 4
My Sqrt(2)  = 1.414213562373095
math.Sqrt(2) = 1.4142135623730951
*/
