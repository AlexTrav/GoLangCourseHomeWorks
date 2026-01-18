package main

import "fmt"

/* ТЗ:
Задание 4. Матрица 3×3 и диагонали
Описание:Создай двумерный слайс matrix := [][]int размером 3×3.
Заполни его числами по формуле matrix[i][j] = i+j.
Напиши две функции:
1. mainDiagonal(matrix [][]int) []int
2. secondaryDiagonal(matrix [][]int) []int
которые возвращают элементы главной и побочной диагонали.
Пример:
[ [0 1 2]  [1 2 3]  [2 3 4] ]
Главная диагональ: [0 2 4]
Побочная диагональ: [2 2 2]
*/

func main() {
	matrix := make([][]int, 3)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, 3)
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = i + j
		}
	}

	fmt.Println("Исходная матрица:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println("Главная диагональ матрицы:", mainDiagonal(matrix))
	fmt.Println("Побочная диагональ матрицы:", secondaryDiagonal(matrix))
}

func mainDiagonal(matrix [][]int) []int {
	diagonal := make([]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		diagonal[i] = matrix[i][i]
	}
	return diagonal
}

func secondaryDiagonal(matrix [][]int) []int {
	diagonal := make([]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		diagonal[i] = matrix[i][len(matrix[i])-i-1]
	}
	return diagonal
}

/* Вывод из консоли:
Исходная матрица:
0 1 2
1 2 3
2 3 4
Главная диагональ матрицы: [0 2 4]
Побочная диагональ матрицы: [2 2 2]
*/
