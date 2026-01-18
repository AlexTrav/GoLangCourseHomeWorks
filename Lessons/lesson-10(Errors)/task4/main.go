package main

import (
	"errors"
	"fmt"
	"strings"
)

/*

## 4) `errors.As` + типизированная ошибка (полезные поля)

**Цель:** делать ошибки, которые можно “распаковать”.

**Задание:** создай тип:

```go
type ValidationErrorstruct {
	Field string
	Reason string
}

```

- Реализуй `Error() string`
- Реализуй `ValidateSignup(email, password string) error`
    - email должен содержать `@`
    - password должен быть минимум 8 символов
    - на первое найденное нарушение возвращай `ValidationError`

**Готово, если:**

- в коде можно сделать:
    - `var ve *ValidationError; errors.As(err, &ve)` → `true`
    - `ve.Field` правильно заполнен

*/

type ValidationError struct {
	Field  string
	Reason string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Reason)
}

func ValidateSignup(email, password string) error {
	if !strings.Contains(email, "@") {
		return &ValidationError{Field: "email", Reason: "must contain @"}
	}
	if len(password) < 8 {
		return &ValidationError{Field: "password", Reason: "must be at least 8 chars"}
	}
	return nil
}

func main() {
	err := ValidateSignup("no-at-symbol", "12345678")
	fmt.Println("err:", err)

	var ve *ValidationError
	fmt.Println("errors.As:", errors.As(err, &ve))
	if ve != nil {
		fmt.Println("field:", ve.Field)
		fmt.Println("reason:", ve.Reason)
	}

	err = ValidateSignup("ab@gmail.com", "123")
	fmt.Println("err:", err)

	fmt.Println("ok:", ValidateSignup("ab@gmail.com", "12345678"))
}

/* Вывод из консоли:
err: email: must contain @
errors.As: true
field: email
reason: must contain @
err: password: must be at least 8 chars
ok: <nil>
*/
