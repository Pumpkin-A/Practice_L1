package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan bool)

	sleepingTime := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), sleepingTime)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if ctx.Err() == context.DeadlineExceeded {
				ch <- true
				close(ch)
				return
			}
		}
	}()

	fmt.Println("sleep for 5 seconds")
	<-ch
	wg.Wait()
	fmt.Println("sleep for 5 seconds is over")
}
