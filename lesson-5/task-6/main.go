package main

import "fmt"

/* ТЗ:
Задание 6. Сгруппировать всех юзеров по городу.
*/

type User struct {
	Name string
	City string
}

// GroupByCity - группирует пользователей по значению поля City
func GroupByCity(users []User) map[string][]User {
	resultMap := make(map[string][]User)
	for _, user := range users {
		resultMap[user.City] = append(resultMap[user.City], user)
	}
	return resultMap
}

func main() {
	tests := [][]User{
		{
			{"Иван", "Москва"},
			{"Петр", "Москва"},
			{"Сергей", "Казань"},
			{"Андрей", "Сочи"},
			{"Анна", "Казань"},
		},
		{
			{"John", "NY"},
			{"Bob", "LA"},
			{"Alice", "NY"},
		},
	}

	for _, t := range tests {
		fmt.Println("Вход:", t)
		result := GroupByCity(t)
		fmt.Println("Группировка по городам:")
		for city, users := range result {
			fmt.Printf("  %s: %v\n", city, users)
		}
		fmt.Println()
	}
}

/* Вывод из консоли:
Группировка по городам:
  Москва: [{Иван Москва} {Петр Москва}]
  Казань: [{Сергей Казань} {Анна Казань}]
  Сочи: [{Андрей Сочи}]

Вход: [{John NY} {Bob LA} {Alice NY}]
Группировка по городам:
  NY: [{John NY} {Alice NY}]
  LA: [{Bob LA}]
*/
