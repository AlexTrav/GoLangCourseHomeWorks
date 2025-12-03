package main

/* ТЗ:
Задача 9. Вложенные интерфейсы (композиция интерфейсов)
Условие:
Создай интерфейсы:
```go
type Reader interface {
	Read() string
}

type Writer interface {
	Write(s string)
}
```
Создай интерфейс `ReadWriter`, который **объединяет** оба.
Реализуй `ReadWriter` в типе `MemoryBuffer`, который хранит последнюю записанную строку.
*/

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(s string)
}

// ReadWriter - композиция интерфейсов
type ReadWriter interface {
	Reader
	Writer
}

type MemoryBuffer struct {
	data string
}

func (m *MemoryBuffer) Write(s string) {
	m.data = s
}

func (m *MemoryBuffer) Read() string {
	return m.data
}

func main() {
	var rw ReadWriter = &MemoryBuffer{}

	rw.Write("test")
	fmt.Println(rw.Read())
}

/* Вывод из консоли:
test
*/
