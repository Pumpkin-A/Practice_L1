// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	i atomic.Int32
}

func (c *Counter) count() {
	c.i.Add(1)
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
	fmt.Println(c.i.Load())
}
