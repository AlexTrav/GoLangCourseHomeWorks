package main

/* ТЗ:
Задача 7. Встраивание структур и методы
Условие:
Создай структуру `User` с полем `Name` и методом `Greet()`.
Создай структуру `Admin`, которая встраивает `User` (анонимное поле).
Покажи, что у `Admin` автоматически есть метод `Greet()`, и добавь ещё метод `Ban()`.
*/

import "fmt"

type User struct {
	Name string
}

func (u User) Greet() {
	fmt.Println("Hello, I am", u.Name)
}

// Admin встраивает User
type Admin struct {
	User
	Level int
}

func (a Admin) Ban(userName string) {
	// a.Name доступен напрямую благодаря встраиванию User
	fmt.Println(a.Name, "bans", userName)
}

func main() {
	a := Admin{User: User{Name: "SuperAdmin"}, Level: 10}

	a.Greet() // метод Greet у User промоутиться в Admin
	a.Ban("troll42")
}

/* Вывод из консоли:
Hello, I am SuperAdmin
SuperAdmin bans troll42
*/
