package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var wg sync.WaitGroup
	exit := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-exit:
				fmt.Println("goroutine close")
				return
			default:
				fmt.Println("goroutine is working")
			}
		}
	}()

	wg.Wait()
	fmt.Println("main close")

}
