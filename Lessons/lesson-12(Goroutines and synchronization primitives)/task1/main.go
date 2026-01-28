package main

/*
## 1) безопасный счётчик

**Цель:** понять, зачем нужен `sync.Mutex`.

**Задание:**
- Создай структуру `Counter` с полем `value int`
- Добавь `Mutex`
- Реализуй метод `Inc()`
- Запусти 1000 горутин, каждая увеличивает счётчик на 1
- В конце выведи значение

**Готово, если:**
- итоговое значение всегда `1000`
*/

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()
	fmt.Println("counter =", counter.value)
}
