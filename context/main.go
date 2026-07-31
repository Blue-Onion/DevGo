package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	exampleContext()
}
func exampleContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	done := make(chan struct{})
	go func() {
		time.Sleep(4 * time.Second)
		close(done)
	}()
	select {
	case <-done:
		fmt.Println("Called The api")
	case <-ctx.Done():
		fmt.Println("Oh timeout expired")
	}
}
