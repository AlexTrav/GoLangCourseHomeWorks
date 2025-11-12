package main

import (
	"fmt"
)

/* ТЗ:
Задание 4. Встраивание (embedding) и promotion полей/методов

Цель: понять, как поля/методы “поднимаются” во внешний тип.
Сделать: `Named{Name string}` и `User{Named; Email string}` (именно встраивание, без имени поля). `Named.String()` возвращает `Name`.
Теория: у `User` доступны `u.Name` и `u.String()` напрямую.
Критерии: `fmt.Sprint(User{Named{"Bob"},"b@x"})` даёт `"Bob"` (в составе).
*/

type Named struct {
	Name string
}

// String - возвращает имя
func (n Named) String() string {
	return n.Name
}

// User - встраивает Named, поэтому поля и методы Named доступны напрямую
type User struct {
	Named
	Email string
}

func main() {
	u := User{
		Named: Named{Name: "Bob"},
		Email: "b@x",
	}

	fmt.Printf("User: %+v\n", u)
	fmt.Println("Имя:", u.Name)
	fmt.Println("Email:", u.Email)
	fmt.Println("String():", u.String()) // доступен напрямую

	// Проверка: (проходит успешно)
	if u.Name != "Bob" {
		panic("promotion failed")
	}
}

/* Вывод из консоли:
User: {Named:{Name:Bob} Email:b@x}
Имя: Bob
Email: b@x
String(): Bob
*/
