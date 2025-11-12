package main

import "fmt"

/* ТЗ:
Задание 7. Глубокое vs поверхностное копирование

Цель: показать, почему простого присваивания недостаточно для ссылочных полей.
Сделать:
`DB{DSN string}`, `Config{Hosts []string; DB *DB}`.
Методы `Shallow() Config` и `Deep() Config`. В `Deep` копируйте слайс `Hosts` и структуру `DB`.
Критерии: изменение копии `Deep` не влияет на оригинал; изменение `Shallow` — влияет.
*/

type DB struct {
	DSN string
}

type Config struct {
	Hosts []string
	DB    *DB
}

// Shallow - поверхностная копия (данные общие)
func (c Config) Shallow() Config {
	return c
}

// Deep - глубокая копия (данные изолированы)
func (c Config) Deep() Config {
	out := c
	// копируем слайс
	if c.Hosts != nil {
		out.Hosts = append([]string(nil), c.Hosts...)
	}
	// копируем структуру DB
	if c.DB != nil {
		cp := *c.DB
		out.DB = &cp
	}
	return out
}

func main() {
	orig := Config{
		Hosts: []string{"a"},
		DB:    &DB{DSN: "x"},
	}

	sh := orig.Shallow()
	dp := orig.Deep()

	fmt.Println("Исходный:", orig)
	fmt.Println("Поверхностная копия:", sh)
	fmt.Println("Глубокая копия:", dp)

	// Меняем поверхностную копию
	sh.Hosts[0] = "b"
	sh.DB.DSN = "y"

	fmt.Println("\nПосле изменения Shallow:")
	fmt.Println("orig:", orig) // изменится
	fmt.Println("sh:", sh)

	// Меняем глубокую копию
	dp.Hosts[0] = "c"
	dp.DB.DSN = "z"

	fmt.Println("\nПосле изменения Deep:")
	fmt.Println("orig:", orig) // не изменится
	fmt.Println("dp:", dp)

	// Проверка: (проходит успешно)
	if orig.Hosts[0] != "b" || orig.DB.DSN != "y" {
		panic("shallow expected to affect orig")
	}
	if orig.Hosts[0] == "c" || orig.DB.DSN == "z" {
		panic("deep must isolate")
	}
}

/* Вывод из консоли:
Исходный: {[a] 0xc0000280b0}
Поверхностная копия: {[a] 0xc0000280b0}
Глубокая копия: {[a] 0xc0000280d0}

После изменения Shallow:
orig: {[b] 0xc0000280b0}
sh: {[b] 0xc0000280b0}

После изменения Deep:
orig: {[b] 0xc0000280b0}
dp: {[c] 0xc0000280d0}
*/
