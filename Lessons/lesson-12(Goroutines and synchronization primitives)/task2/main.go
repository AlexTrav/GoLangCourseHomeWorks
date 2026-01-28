package main

/*
## 2) RWMutex

**Цель:** показать разницу между `Mutex` и `RWMutex`.

**Задание:**
- Структура `Storage` с `map[string]int`
- Используй `sync.RWMutex`
- Методы:
  - `Set(key string, value int)`
  - `Get(key string) int`
- Запусти:
  - 100 горутин `Get`
  - 10 горутин `Set`

**Готово, если:**
- нет data race
*/

import (
	"fmt"
	"sync"
)

type Storage struct {
	mu   sync.RWMutex
	data map[string]int
}

func (s *Storage) Set(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *Storage) Get(key string) int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data[key]
}

func main() {
	var wg sync.WaitGroup
	store := Storage{data: make(map[string]int)}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			store.Set("a", i)
		}(i)
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(store.Get("a"))
		}()
	}

	wg.Wait()
}
