package main

/*
## 4) семафор через Mutex

**Цель:** контроль параллелизма без каналов.

**Задание:**
- Не более 5 горутин могут работать одновременно
- Используй `Mutex` + счётчик
- 20 задач
- Каждая задача "работает" 100мс

**Готово, если:**
- одновременно выполняется максимум 5 задач
*/

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	mu    sync.Mutex
	limit int
	count int
}

func (s *Semaphore) Acquire() {
	for {
		s.mu.Lock()
		if s.count < s.limit {
			s.count++
			s.mu.Unlock()
			return
		}
		s.mu.Unlock()
		time.Sleep(10 * time.Millisecond)
	}
}

func (s *Semaphore) Release() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count--
}

func main() {
	var wg sync.WaitGroup
	sem := Semaphore{limit: 5}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			sem.Acquire()
			fmt.Println("start", id)
			time.Sleep(100 * time.Millisecond)
			fmt.Println("end", id)
			sem.Release()
		}(i)
	}

	wg.Wait()
}
