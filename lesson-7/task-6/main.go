package main

/*
Задача 6. Указатель/значение и интерфейс
Условие:
Создай интерфейс `Updater` с методом `Update(msg string)`.
Создай тип `Logger` с полем `Last string`.
Сделай метод `Update` с **указательным** получателем, который сохраняет `msg` в `Last`.
Попробуй присвоить в переменную типа `Updater` и `Logger`, и `*Logger`. Посмотри, что компилируется.
Решение:
*/

import "fmt"

type Updater interface {
	Update(msg string)
}

type Logger struct {
	Last string
}

// Update - Метод с указательным получателем
func (l *Logger) Update(msg string) {
	l.Last = msg
}

func main() {
	var u1 Updater
	_ = u1

	var u2 Updater

	l := Logger{}

	// u1 = l // <- так НЕ компилируется, потому что l - значение Logger
	u2 = &l // <- а так компилируется, потому что &l - *Logger

	u2.Update("hello")
	fmt.Println("Last:", l.Last)
}

/* Вывод из консоли:
Last: hello
*/
