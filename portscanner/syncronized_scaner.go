package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	// DEclare a waitgroup which is an struct
	var wg sync.WaitGroup
	for i := 0; i <= 65000; i++ {
		// Add a waitgroup - it willincrement the counter with the value supplied
		wg.Add(1)
		go func(j int) {
			// fmt.Println("Passed j ", j)
			defer wg.Done()
			address := fmt.Sprintf("localhost:%d", j)
			// fmt.Println(address)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				// fmt.Println("[!]", err)
				return
			}
			conn.Close()
			fmt.Println("[+] OPEN -> ", j)
		}(i)
	}
	wg.Wait()
}
