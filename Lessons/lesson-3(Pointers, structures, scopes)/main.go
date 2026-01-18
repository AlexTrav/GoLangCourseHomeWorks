package main

import "fmt"

/* ТЗ:
Создай type Person struct { Name string; Age int }.
Напиши NewPerson(name string) и метод Birthday() (+1 к возрасту). Проверь разницу между значением и указателем:
p.Birthday() для p Person vs *Person.

Сделай Book{Title, Author, Pages}. Создай через литерал, через конструктор, через new(Book). Выведи значения - что по умолчанию?

Вложенные структуры.
User{Name, Address} и Address{City, Street}. Напиши функцию FullAddress(u User) string.
*/

// Первое задание:

type Person struct {
	Name string
	Age  int
}

// NewPerson - Конструктор: возвращает указатель, чтобы можно было изменять объект напрямую
func NewPerson(name string, age int) *Person {
	return &Person{name, age}
}

// Birthday - метод с указателем (меняет оригинал)
func (p *Person) Birthday() {
	p.Age++
}

// BirthdayValue - метод с копией (не меняет оригинал)
func (p Person) BirthdayValue() {
	p.Age++
}

// Второе задание:

type Book struct {
	Title  string
	Author string
	Pages  int
}

// NewBook - Конструктор (возвращает указатель на Book)
func NewBook(title, author string, pages int) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}

// Третье задание:

type Address struct {
	City   string
	Street string
}

type User struct {
	Name    string
	Address Address
}

// FullAddress - возвращает строку с полным адресом
func (u User) FullAddress() string {
	return fmt.Sprintf("%s, %s", u.Address.City, u.Address.Street)
}

func main() {
	// Первое задание:
	fmt.Println("Задание 1 - Результат: ")

	// Через указатель
	alex := NewPerson("Алексей", 22)
	alex.Birthday()                                 // изменит оригинал
	fmt.Println("После alex.Birthday():", alex.Age) // 23

	// Через значение
	bob := Person{"Боб", 30}
	bob.BirthdayValue()                                // не изменит оригинал
	fmt.Println("После bob.BirthdayValue():", bob.Age) // 30

	// Но если вызвать метод с указателем на значение - Go сам создаст указатель
	bob.Birthday()
	fmt.Println("После bob.Birthday():", bob.Age) // 31

	// Разница:	(p Person) создаёт копию и не изменяет оригинал,
	// а (*p Person) работает с указателем и изменяет исходную структуру напрямую.

	// Второе задание:
	fmt.Println("\nЗадание 2 - Результат: ")

	// Создание через литерал
	book1 := Book{"1984", "Джордж Оруэлл", 328}
	fmt.Println("book1:", book1)

	// Создание через конструктор
	book2 := NewBook("О дивный новый мир", "Олдос Хаксли", 288)
	fmt.Println("book2:", *book2) // разыменовываем, чтобы показать значения

	// Создание через new(Book)
	book3 := new(Book)
	fmt.Println("book3:", *book3) // поля по умолчанию

	// Изменил значения book3
	book3.Title = "451 Градус по Фаренгейту"
	book3.Author = "Рэй Брэдбери"
	book3.Pages = 256
	fmt.Println("После заполнения book3:", *book3)

	// По умолчанию: строки - "", числа - 0, bool - false, указатели - nil.
	// Разница: Book{} создаёт значение, new(Book) - указатель с нулевыми полями, конструктор задаёт свои значения.

	// Третье задание:
	fmt.Println("\nЗадание 3 - Результат:")

	user := User{
		Name: "Алиса",
		Address: Address{
			City:   "Алматы",
			Street: "Проспект Абая, 25а",
		},
	}

	fmt.Printf("Полный адрес пользователя %s: %s", user.Name, user.FullAddress())
}

/* Вывод из консоли:
Задание 1 - Результат:
После alex.Birthday(): 23
После bob.BirthdayValue(): 30
После bob.Birthday(): 31

Задание 2 - Результат:
book1: {1984 Джордж Оруэлл 328}
book2: {О дивный новый мир
 Олдос Хаксли 288}
book3: {  0}
После заполнения book3: {451 Градус по Фаренгейту Рэй Брэдбери 256}

Задание 3 - Результат:
Полный адрес пользователя Алиса: Алматы, Проспект Абая, 25а
*/
