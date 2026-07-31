package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"time"
)

func main() {
	done := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go func() {
		if rand.Int() > 3 {
			slog.Info("Bad Luck")

			select {
			case <-time.After(100 * time.Second):
				fmt.Println("Finished work")
			case <-ctx.Done():
				fmt.Println("Cancelled!")
				return
			}
			return
		}

		select {
		case done <- 1:
		case <-ctx.Done():
			return
		}
	}()

	select {
	case <-done:
		fmt.Println("Success")
	case <-ctx.Done():
		fmt.Println("Timeout")
	}
}
