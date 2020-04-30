// SImple IMplementation of GOlang
package main

import "fmt"

func main() {

	stack := []string{}
	var choice string
	var value string
	for {
		fmt.Println("CURRENT STACK IS :-", stack)
		fmt.Print("APPEND / POP (A/P) > ")
		fmt.Scan(&choice)

		switch choice {
		case "A":
			fmt.Print("ENter the string to append >")
			fmt.Scan(&value)
			stack = append(stack, value)
		case "P":
			if len(stack) == 0 {
				fmt.Println("[!] Stack is EMpty")
			} else {
				n := len(stack) - 1
				stack = stack[:n]
			}
		}
	}
}
