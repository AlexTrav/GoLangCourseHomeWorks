package main

/* ТЗ:
Задача 3. Простой интерфейс и одна реализация
Условие:
Создай интерфейс `Greeter` с методом `Greet() string`.
Создай тип `Person` с полем `Name string`, который реализует этот интерфейс (метод возвращает строку `"Hello, <Name>"`).
Напиши функцию `SayHello(g Greeter)`, которая печатает результат `Greet()`.
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

func SayHello(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {
	p := Person{Name: "Фёдор"}
	SayHello(p)
}

/* Вывод из консоли:
Hello, Фёдор
*/
