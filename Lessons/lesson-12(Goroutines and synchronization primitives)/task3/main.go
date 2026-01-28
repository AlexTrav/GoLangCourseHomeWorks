package main

/*
## 3) банковский счёт

**Цель:** атомарность операций.

**Задание:**
- Структура `Account`
- Поля: `balance int`, `Mutex`
- Методы:
  - `Deposit(sum int)`
  - `Withdraw(sum int) bool`
- 1000 горутин кладут по 10
- 1000 горутин снимают по 10

**Готово, если:**
- баланс всегда корректный
*/

import (
	"fmt"
	"sync"
)

type Account struct {
	mu      sync.Mutex
	balance int
}

func (a *Account) Deposit(sum int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += sum
}

func (a *Account) Withdraw(sum int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.balance < sum {
		return false
	}
	a.balance -= sum
	return true
}

func main() {
	var wg sync.WaitGroup
	acc := Account{balance: 0}

	for i := 0; i < 1000; i++ {
		wg.Add(2)

		go func() {
			defer wg.Done()
			acc.Deposit(10)
		}()

		go func() {
			defer wg.Done()
			acc.Withdraw(10)
		}()
	}

	wg.Wait()
	fmt.Println("balance =", acc.balance)
}
