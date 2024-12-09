package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ch := make(chan bool)

	sleepingTime := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), sleepingTime)
	g, gCtx := errgroup.WithContext(ctx)
	defer cancel()

	g.Go(func() error {
		for {
			if gCtx.Err() == context.DeadlineExceeded {
				ch <- true
				close(ch)
				return nil
			}
		}
	})

	fmt.Println("sleep for 5 seconds")
	<-ch
	fmt.Println("sleep for 5 seconds is over")
}
