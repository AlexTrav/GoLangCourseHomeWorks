package main

/* ТЗ:
Задача 5. Срез интерфейсов
Условие:
Используя интерфейс `Greeter` и типы `Person` и `Robot`, сделай срез `[]Greeter` и в цикле вызови `Greet()` у каждого элемента.
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

func main() {
	var greeters []Greeter

	// инициализируем слайс разными реализациями интерфейса
	greeters = []Greeter{
		Person{Name: "Фёдор"},
		Person{Name: "Анна"},
		Robot{ID: "RX-78"},
		Robot{ID: "T-800"},
	}

	for _, g := range greeters {
		fmt.Println(g.Greet())
	}
}

/* Вывод из консоли:
Hello, Фёдор
Hello, Анна
Beep. I am robot RX-78
Beep. I am robot T-800
*/
