// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	i  int
	mu sync.Mutex
}

func (c *Counter) count() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.i++
}

func main() {
	c := &Counter{}
	var wg sync.WaitGroup
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.count()
		}()
	}
	wg.Wait()
	fmt.Println(c.i)
}
