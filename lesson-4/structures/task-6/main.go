package main

import "fmt"

/* ТЗ:
Цель: прочувствовать, как набор методов влияет на реализацию интерфейса.
Сделать: `Counter{n int}` с `Add(k int)` (получатель-значение) и `Inc()` (получатель-указатель). Интерфейс `Incer{ Inc() }`.
Критерии: `*Counter` реализует `Incer`, `Counter` — нет.
*/

type Counter struct {
	n int
}

// Add - метод со значением (работает с копией)
func (c Counter) Add(k int) {
	c.n += k
	fmt.Println("Внутри Add, n стало:", c.n)
}

// Inc - метод с указателем (меняет оригинал)
func (c *Counter) Inc() {
	c.n++
	fmt.Println("Внутри Inc, n стало:", c.n)
}

// Incer - Интерфейс, требующий только Inc()
type Incer interface {
	Inc()
}

func wantIncer(x Incer) {
	fmt.Println("Тип реализует интерфейс Incer")
}

func main() {
	var c Counter       // значение
	var pc = &Counter{} // указатель

	fmt.Println("== Проверка Add (значение) ==")
	c.Add(5)
	fmt.Println("После Add, оригинал c:", c.n) // останется 0

	fmt.Println("\n== Проверка Inc (указатель) ==")
	pc.Inc() // изменяет оригинал
	fmt.Println("После Inc, pc.n:", pc.n)

	fmt.Println("\n== Проверка интерфейсов ==")

	// wantIncer(c) // Не скомпилируется: Counter не реализует Incer
	wantIncer(pc) // ОК: *Counter реализует Incer
}

/* Вывод из консоли:
== Проверка Add (значение) ==
Внутри Add, n стало: 5
После Add, оригинал c: 0

== Проверка Inc (указатель) ==
Внутри Inc, n стало: 1
После Inc, pc.n: 1

== Проверка интерфейсов ==
Тип реализует интерфейс Incer
*/
