// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"fmt"
	"sync"
	"time"
)

var flag bool

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(flag *bool) {
		defer wg.Done()
		for {
			fmt.Println("goroutine os working")
			if *flag {
				fmt.Println("goroutine close")
				return
			}
		}
	}(&flag)

	time.Sleep(3 * time.Second)
	flag = true
	wg.Wait()
	fmt.Println("main close")

}
