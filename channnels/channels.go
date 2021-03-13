package main

import "fmt"

// func greet(c chan string) {
// 	msg := <-c
// 	fmt.Println("Data recieved from channel ", msg)
// }

func squares(c chan string) {
	// for i := 0; i <= 9; i++ {
	// 	c <- i * i
	// }
	// fmt.Println("val recieved ", <-c)
	s := <-c
	fmt.Println(s)
	// words := []string{"Yo9", "Yo8", "Yo7", "Yo6", "Yo5", "Yo4", "Yo3", "Yo2", "Yo1"}
	// for _, word := range words {
	// 	c <- word
	// }

	// close(c)
}

func main() {
	c := make(chan string, 4)
	fmt.Println(" type pf c ", c)
	go squares(c)
	// fs
	c <- "as"
	c <- "as1"
	c <- "as2"
	c <- "as3"
	c <- "as4"

	fmt.Println("done")

}
