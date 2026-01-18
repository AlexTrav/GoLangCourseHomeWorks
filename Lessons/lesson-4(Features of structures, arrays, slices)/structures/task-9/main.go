package main

import (
	"fmt"
	"time"
)

/* ТЗ:
Задание 9. Базовая модель + “хуки” через embedding

Цель: повторное использование полей/логики через встраивание, имитация audit-полей.
Сделать:
`Timestamps{CreatedAt, UpdatedAt time.Time}` + метод `Touch()` (обновляет `UpdatedAt`).
`Order{Timestamps; ID string; Items []string}`.
`NewOrder(id string, items ...string) *Order` заполняет обе даты текущим временем.
Метод `AddItem(it string)` (указатель!) добавляет товар и вызывает `Touch()`.
Подводные камни: используйте указатели для мутаций; в тестах лучше сравнивать, что `UpdatedAt` увеличивается.
Критерии: `UpdatedAt` меняется только при мутации.
*/

// Timestamps - Базовые audit-поля с методом обновления времени
type Timestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Touch - обновляет время обновления
func (t *Timestamps) Touch() {
	t.UpdatedAt = time.Now()
}

// Order - встраивает Timestamps и содержит данные заказа
type Order struct {
	Timestamps
	ID    string
	Items []string
}

// NewOrder - создаёт новый заказ с текущими датами
func NewOrder(id string, items ...string) *Order {
	now := time.Now()
	return &Order{
		Timestamps: Timestamps{
			CreatedAt: now,
			UpdatedAt: now,
		},
		ID:    id,
		Items: append([]string(nil), items...),
	}
}

// AddItem - добавляет товар и обновляет UpdatedAt
func (o *Order) AddItem(it string) {
	o.Items = append(o.Items, it)
	o.Touch()
}

func main() {
	o := NewOrder("o1", "apple")
	fmt.Printf("Создан заказ: %+v\n", o)

	time.Sleep(10 * time.Millisecond) // имитируем задержку
	old := o.UpdatedAt

	o.AddItem("banana")
	fmt.Printf("После AddItem: %+v\n", o)

	// Проверка: (проходит успешно)
	if !o.UpdatedAt.After(old) {
		panic("Touch() not called")
	}

}

/* Вывод из консоли:
Создан заказ: &{Timestamps:{CreatedAt:2025-11-12 18:36:56.7338848 +0500 +05 m=+0.000506801 UpdatedAt:2025-11-12 18:36:56.7338848 +0500 +05 m=+0.000506801} ID:o1 Items:[apple]}
После AddItem: &{Timestamps:{CreatedAt:2025-11-12 18:36:56.7338848 +0500 +05 m=+0.000506801 UpdatedAt:2025-11-12 18:36:56.7716533 +0500 +05 m=+0.038275301} ID:o1 Items:[apple banana]}
*/
