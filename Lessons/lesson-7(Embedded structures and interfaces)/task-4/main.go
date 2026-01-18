package main

/* ТЗ:
Задача 4. Полиморфизм: несколько реализаций одного интерфейса
Условие:
Добавь тип `Robot` с полем `ID string`, который тоже реализует `Greeter`, но приветствует так: `"Beep. I am robot <ID>"`.
В `main` передай в `SayHello` и `Person`, и `Robot`.
*/

import "fmt"

type Greeter interface {
	Greet() string
}

type Person struct {
	Name string
}

func (p Person) Greet() string {
	return "Hello, " + p.Name
}

type Robot struct {
	ID string
}

func (r Robot) Greet() string {
	return "Beep. I am robot " + r.ID
}

func SayHello(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {
	p := Person{Name: "Фёдор"}
	r := Robot{ID: "RX-78"}

	SayHello(p)
	SayHello(r)
}

/* Вывод из консоли:
Hello, Фёдор
Beep. I am robot RX-78
*/
