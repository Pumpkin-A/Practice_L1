// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ch := make(chan int)

	workingTime := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), workingTime)
	g, gCtx := errgroup.WithContext(ctx)
	defer cancel()

	g.Go(func() error {
		var i int
		for {
			if gCtx.Err() == context.DeadlineExceeded {
				close(ch)
				return nil
			}
			ch <- i
			i++
		}
	})

	g.Go(func() error {
		for {
			v, ok := <-ch
			if !ok {
				return nil
			}
			fmt.Println(v)
		}
	})

	err := g.Wait()
	if err != nil {
		fmt.Println("Error group: ", err)
	}
	fmt.Println("Main done")

}
