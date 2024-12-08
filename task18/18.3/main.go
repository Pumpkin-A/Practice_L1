// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	i  int
	ch chan struct{}
}

func (c *Counter) count() chan struct{} {
	chFinish := make(chan struct{})

	go func() {
		defer func() {
			chFinish <- struct{}{}
			close(chFinish)
		}()
		for {
			_, ok := <-c.ch
			if !ok {
				return
			}
			c.i++
		}
	}()

	return chFinish
}

func main() {
	var wg sync.WaitGroup
	c := &Counter{ch: make(chan struct{})}

	chFinish := c.count()

	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.ch <- struct{}{}
		}()
	}

	wg.Wait()
	close(c.ch)
	<-chFinish

	fmt.Println(c.i)
}
