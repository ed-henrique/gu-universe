package main

import (
	"fmt"
	"gu-universe/internal/client"
	"sync"
)

func main() {
	clients := make([]client.Client, 10_000)
	for i := range len(clients) {
		clients[i] = client.New()
	}

	var (
		errChannel = make(chan error, 10_000)
		wg         = &sync.WaitGroup{}
	)

	for _, c := range clients {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := c.Start("localhost:8080"); err != nil {
				errChannel <- err
				return
			}
		}()
	}

	go func() {
		wg.Add(1)
		defer wg.Done()

		select {
		case err := <-errChannel:
			fmt.Println(err)
		}
	}()

	wg.Wait()
	close(errChannel)
}
