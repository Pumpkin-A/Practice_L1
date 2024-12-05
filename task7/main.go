// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)
	mapa := make(map[int]int)
	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 100 {
				mu.Lock()
				mapa[rand.Intn(100)] = rand.Int()
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	for k, v := range mapa {
		fmt.Println(k, ":", v)
	}
}
