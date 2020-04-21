package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("localhost:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}

}

func main() {
	// 100 is the buffer
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		// generate threadpool
		go worker(ports, results)
	}

	go func() {
		// Send ports through ports channel
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 1; i < 1500; i++ {
		port := <-results
		if port != 0 {
			fmt.Println("[+] FOund Open Port :-", port)
			openports = append(openports, port)
		}
	}

	// CLose Channels
	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Println("OPEN -> ", port)
	}
}
