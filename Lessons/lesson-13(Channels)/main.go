package main

/* Домашка

есть магазин apple Store
у него должен быть счет в банке

добавить сущность покупателя хипстера со своим счетом

вышел новый айфон

создаем очередь из 10 000 хипстеров со своими кошельками и балансами у входа в магазин

делаем обратный отсчет до начала продаж

магазин начинает работу, но может обслужить одновременно не более 300 человек

каждый параллельно пытается купить айфон, если денег хватает, делает перевод со своего счета на счёт магазина,
и забирает айфон и уходит, если нет, то просто уходит и плачет в старбаксе

в конце вывести статистику, сколько человек купило телефон
сколько не купили из-за недостатка средств
Сколько в общем прошло людей

+(Сделал дополнительно с браком и вовзратами)

*/

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Банковский счёт с mutex
type BankAccount struct {
	Balance int
	mu      sync.Mutex
}

// Пополнение счёта
func (a *BankAccount) Deposit(amount int) {
	a.mu.Lock()
	a.Balance += amount
	a.mu.Unlock()
}

// Снятие денег, если хватает
func (a *BankAccount) Withdraw(amount int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.Balance >= amount {
		a.Balance -= amount
		return true
	}
	return false
}

// Покупатель
type Hipster struct {
	ID     int
	Wallet *BankAccount
}

// Магазин
type Store struct {
	Account *BankAccount
	Price   int // цена айфона
}

// Статистика
var (
	bought   int64 // купили
	failed   int64 // не хватило денег
	total    int64 // всего пришло
	returned int64 // возвраты
)

// Воркер покупки (касса)
func buyWorker(store *Store, buyQueue <-chan Hipster, returnQueue chan<- Hipster, wg *sync.WaitGroup) {
	defer wg.Done()

	for h := range buyQueue {
		atomic.AddInt64(&total, 1) // увеличиваем счётчик людей

		// Пытаемся купить
		if h.Wallet.Withdraw(store.Price) {
			store.Account.Deposit(store.Price) // деньги магазину
			atomic.AddInt64(&bought, 1)

			// 5% шанс возврата
			if rand.Intn(100) < 5 {
				returnQueue <- h
			}
		} else {
			atomic.AddInt64(&failed, 1) // денег нет
		}
	}
}

// Воркер возврата
func returnWorker(store *Store, returnQueue <-chan Hipster, wg *sync.WaitGroup) {
	defer wg.Done()

	for h := range returnQueue {
		// Возвращаем деньги
		if store.Account.Withdraw(store.Price) {
			h.Wallet.Deposit(store.Price)
			atomic.AddInt64(&returned, 1)
		}
	}
}

func main() {
	// Создаём магазин
	store := &Store{
		Account: &BankAccount{Balance: 0},
		Price:   1000,
	}

	// Очереди
	buyQueue := make(chan Hipster, 10000)
	returnQueue := make(chan Hipster, 1000)

	// Таймер открытия
	fmt.Println("Магазин откроется через 3 секунды...")
	timer := time.NewTimer(3 * time.Second)
	<-timer.C
	fmt.Println("Магазин открыт!")

	// Генерация 10 000 покупателей
	for i := 0; i < 10000; i++ {
		balance := rand.Intn(2000) // случайный баланс
		h := Hipster{
			ID:     i,
			Wallet: &BankAccount{Balance: balance},
		}
		buyQueue <- h
	}
	close(buyQueue) // больше не добавляем

	// Запуск 300 касс
	var buyWG sync.WaitGroup
	for i := 0; i < 300; i++ {
		buyWG.Add(1)
		go buyWorker(store, buyQueue, returnQueue, &buyWG)
	}

	buyWG.Wait()       // ждём покупки
	close(returnQueue) // закрываем возвраты

	// Запуск 10 воркеров возврата
	var returnWG sync.WaitGroup
	for i := 0; i < 10; i++ {
		returnWG.Add(1)
		go returnWorker(store, returnQueue, &returnWG)
	}

	returnWG.Wait() // ждём возвраты

	// Итоговая статистика
	fmt.Println("----- СТАТИСТИКА -----")
	fmt.Println("Всего людей прошло:", total)
	fmt.Println("Купили телефон:", bought)
	fmt.Println("Не хватило денег:", failed)
	fmt.Println("Вернули телефон:", returned)
	fmt.Println("Баланс магазина:", store.Account.Balance)
}

/* Вывод из консоли:
Магазин откроется через 3 секунды...
Магазин открыт!
----- СТАТИСТИКА -----
Всего людей прошло: 10000
Купили телефон: 4997
Не хватило денег: 5003
Вернули телефон: 273
Баланс магазина: 4724000
*/
