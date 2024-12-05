// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
// которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора
// количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	var workersNumber int
	fmt.Print("Print number of workers: ")
	_, err := fmt.Scanf("%d", &workersNumber)
	if err != nil {
		fmt.Println("int expected")
		return
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	for i := range workersNumber {
		wg.Add(1)
		go func(ctx context.Context, i int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("nnn")
					return
				default:
					fmt.Println(i, <-ch)
				}
			}
		}(ctx, i)
	}

	go func() {
		<-ctx.Done()
		close(ch)
		wg.Wait()
		return
	}()

	var i int
	for {

		fmt.Println("отправляем:", i)
		ch <- i
		i++
	}

	fmt.Println("все хорошо")

}
