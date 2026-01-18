package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

/*

## 2) Контекст + обёртки

**Цель:** правильно добавлять контекст через `%w` и не терять причину.

**Задание:** реализуй:

- `ReadFileTrim(path string) (string, error)`
    - читает файл `os.ReadFile`
    - если ошибка → `fmt.Errorf("read %q: %w", path, err)`
    - возвращает строку `strings.TrimSpace(...)`

**Проверки:**

- при отсутствующем файле `errors.As(err, *os.PathError)` должно быть `true`
- сообщение должно содержать `read "..."`

*/

func main() {
	s, err := ReadFileTrim("tmp.txt")
	if err != nil {
		// обработка ошибки + проверка, что причина не потерялась
		var pe *os.PathError
		if errors.As(err, &pe) {
			fmt.Println("file error:", err)
			return
		}
		fmt.Println("other error:", err)
		return
	}

	fmt.Println(s)
}

func ReadFileTrim(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read %q: %w", path, err)
	}
	return strings.TrimSpace(string(b)), nil
}

/* Вывод из консоли:
file error: read "tmp.txt": open tmp.txt: The system cannot find the file specified.
*/
