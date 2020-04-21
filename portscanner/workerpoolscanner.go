package main

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	// accept channel of int, pointer of waitgroup
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}

func main() {
	// CHannel of ports
	ports := make(chan int, 10)
	var wg sync.WaitGroup
	for i := 0; i < cap(ports); i++ {
		// Loop to create resource pool
		go worker(ports, &wg)
	}
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}
