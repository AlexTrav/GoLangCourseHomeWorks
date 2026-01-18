package main

import (
	"errors"
	"fmt"
)

/*

## 3) `errors.Is` и “сентинел”: поиск пользователя

**Цель:** научиться отличать “не найдено” от “прочих”.

**Дано:**

```go
var ErrNotFound = errors.New("not found")

```

**Задание:** реализуй:

- `FindUserID(name string) (int, error)`
    - если `name == ""` → `fmt.Errorf("empty name: %w", ErrNotFound)` (да, тут специально)
    - если `name == "admin"` → `return 1, nil`
    - иначе → `return 0, fmt.Errorf("user %q: %w", name, ErrNotFound)`
- `IsNotFound(err error) bool` через `errors.Is(err, ErrNotFound)`

**Готово, если:**

- `IsNotFound(err)` true для пустого имени и неизвестного имени
- `admin` не даёт ошибку

*/

var ErrNotFound = errors.New("not found")

func main() {
	id, err := FindUserID("")
	fmt.Println(id, err, "IsNotFound:", IsNotFound(err))

	id, err = FindUserID("bob")
	fmt.Println(id, err, "IsNotFound:", IsNotFound(err))

	id, err = FindUserID("admin")
	fmt.Println(id, err, "IsNotFound:", IsNotFound(err))
}

func FindUserID(name string) (int, error) {
	if name == "" {
		return 0, fmt.Errorf("empty name: %w", ErrNotFound)
	}
	if name == "admin" {
		return 1, nil
	}
	return 0, fmt.Errorf("user %q: %w", name, ErrNotFound)
}

func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

/* Вывод из консоли:
0 empty name: not found IsNotFound: true
0 user "bob": not found IsNotFound: true
1 <nil> IsNotFound: false
*/
