package main

/* ТЗ:
Задача 8. Интерфейс, реализуемый через встраивание
Условие:
Есть интерфейс:
```go
type Greeter interface {
	Greet()
}
```
Сделай так, чтобы `Admin` из предыдущей задачи тоже реализовывал `Greeter`, **ничего не меняя в самом интерфейсе**.
Покажи, что функция `SayGreet(g Greeter)` принимает и `User`, и `Admin`.
*/

import "fmt"

type Greeter interface {
	Greet()
}

type User struct {
	Name string
}

func (u User) Greet() {
	fmt.Println("Hello, I am", u.Name)
}

// Admin встраивает User, поэтому наследует его методы
type Admin struct {
	User
	Level int
}

func SayGreet(g Greeter) {
	g.Greet()
}

func main() {
	u := User{Name: "User1"}
	a := Admin{User: User{Name: "Admin1"}, Level: 10}

	SayGreet(u)
	SayGreet(a) // Admin реализует Greeter через встроенный User
}

/* Вывод из консоли:
Hello, I am User1
Hello, I am Admin1
*/
