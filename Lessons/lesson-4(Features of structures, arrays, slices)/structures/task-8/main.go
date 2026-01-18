package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

/*	ТЗ:

Задание 8. Композиция + JSON-теги + скрытие чувствительных полей

Цель: комбинировать встраивание и управление сериализацией.
Сделать:
`Credentials{User string \`json:"user"`; Password string `json:"-" ` }`.
`Account{ID int `json:"id"`; Credentials}`(embedding).
Реализовать`MarshalJSON`у`Account`, чтобы вернуть` {"id":..., "user":...}`(без пароля).
Подводные камни: избегайте рекурсии в`MarshalJSON` (используйте alias).
*/

type Credentials struct {
	User     string `json:"user"`
	Password string `json:"-"` // Не сериализуется стандартно
}

type Account struct {
	ID int `json:"id"`
	Credentials
}

// MarshalJSON - возвращает JSON без поля Password
func (a Account) MarshalJSON() ([]byte, error) {
	type Alias struct {
		ID   int    `json:"id"`
		User string `json:"user"`
	}
	return json.Marshal(Alias{
		ID:   a.ID,
		User: a.User, // поле User из встроенного Credentials
	})
}

func main() {
	a := Account{
		ID: 1,
		Credentials: Credentials{
			User:     "bob",
			Password: "secret",
		},
	}

	data, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}

	fmt.Println("Результат сериализации JSON:")
	fmt.Println(string(data))

	// Проверка: (проходит успешно)
	b, _ := json.Marshal(Account{ID: 2, Credentials: Credentials{User: "bob", Password: "secret"}})
	if !bytes.Contains(b, []byte(`"user"`)) || bytes.Contains(b, []byte("secret")) {
		panic("password must be hidden")
	}
}

/* Вывод из консоли:
Результат сериализации JSON:
{"id":1,"user":"bob"}
*/
