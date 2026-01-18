package main

import "fmt"

/* ТЗ:
Задание 3. Проверить, является ли слово Анаграммой.
Например, на вход имеем (”кабан”, "банка”) возвращаем true
*/

// IsAnagram - проверяет, являются ли строки a и b анаграммами
func IsAnagram(a, b string) bool {
	if len([]rune(a)) != len([]rune(b)) {
		return false
	}

	mapA := make(map[rune]int)
	mapB := make(map[rune]int)

	for _, v := range a {
		mapA[v]++
	}

	for _, v := range b {
		mapB[v]++
	}

	if len(mapA) != len(mapB) {
		return false
	}

	for k, valA := range mapA {
		if mapB[k] != valA {
			return false
		}
	}

	return true
}

func main() {
	tests := [][2]string{
		{"кабан", "банка"},
		{"кабан", "банки"},
		{"abc", "cba"},
	}

	for _, t := range tests {
		a := t[0]
		b := t[1]

		result := IsAnagram(a, b)

		if result {
			fmt.Printf("Слова \"%s\" и \"%s\" - анаграммы +\n", a, b)
		} else {
			fmt.Printf("Слова \"%s\" и \"%s\" - НЕ анаграммы -\n", a, b)
		}
	}
}

/* Вывод из консоли:
Слова "кабан" и "банка" - анаграммы +
Слова "кабан" и "банки" - НЕ анаграммы -
Слова "abc" и "cba" - анаграммы +
*/
