package main

import "fmt"

/* ТЗ:
Задание 5. Конфликты имён при встраивании и явная квалификация

Цель: научиться разрешать конфликты методов с одинаковыми именами.
Сделать: `Logger.Log(msg)`, `Audit.Log(msg)`. `Service{Logger; Audit}`.
Критерии: вызов `s.Logger.Log()` и `s.Audit.Log()` работает; через `s.Log()` нельзя, т.к. конфликт.
*/

type Logger struct{}

func (Logger) Log(msg string) {
	fmt.Println("[LOG]", msg)
}

type Audit struct{}

func (Audit) Log(msg string) {
	fmt.Println("[AUDIT]", msg)
}

type Service struct {
	Logger
	Audit
}

func doLog(l *Logger, msg string) {
	l.Log(msg)
}

func doAudit(a *Audit, msg string) {
	a.Log(msg)
}

func main() {
	s := Service{}

	// Корректные вызовы с явной квалификацией:
	s.Logger.Log("запуск сервиса")
	s.Audit.Log("аудит события #42")

	// Некорректно - конфликт, ошибка компиляции:
	// s.Log("это не скомпилируется: ambiguous selector")

	// Можно хранить и использовать оба независимо
	doLog(&s.Logger, "сообщение для лога")
	doAudit(&s.Audit, "сообщение для аудита")
}

/* Вывод из консоли:
[LOG] запуск сервиса
[AUDIT] аудит события #42
[LOG] сообщение для лога
[AUDIT] сообщение для аудита
*/
