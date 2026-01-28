package main

/*
## 5) deadlock

**Цель:** увидеть и понять взаимную блокировку.

**Задание:**
- Два `Mutex`
- Две горутины
- Каждая блокирует их в разном порядке
- Исправь deadlock

*/

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu1 sync.Mutex
	var mu2 sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		mu1.Lock()
		defer mu1.Unlock()

		time.Sleep(100 * time.Millisecond)

		mu2.Lock()
		defer mu2.Unlock()

		fmt.Println("goroutine 1 done")
	}()

	go func() {
		defer wg.Done()
		// было: mu2 -> mu1
		// исправлено: mu1 -> mu2
		mu1.Lock()
		defer mu1.Unlock()

		mu2.Lock()
		defer mu2.Unlock()

		fmt.Println("goroutine 2 done")
	}()

	wg.Wait()
}
