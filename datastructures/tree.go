package main

import (
	"fmt"
)

// Node of trees
type Node struct {
	// Every node will have a value, right and left child are other nodes
	Value int
	Left  *Node
	Right *Node
}

func printNode(n *Node) {
	fmt.Println("Value is :- ", n.Value)
	// Check if there is a left node
	if n.Left != nil {
		fmt.Println("Left :- ", n.Left.Value)
	}
	if n.Right != nil {
		fmt.Println("Right :- ", n.Right.Value)
	}

}

func read() []Node {
	// Take input from user
	var N int
	fmt.Print("[+] ENter the Number of NOdes >")
	fmt.Scanf("%d", &N)
	// Create slices of nodes of size N
	var nodes []Node = make([]Node, N)
	var val, leftNode, rightNode int
	for i := 0; i < N; i++ {
		fmt.Print("Enter the Value of the Element ", i, " > ")
		fmt.Scan(&val)
		nodes[i].Value = val
		fmt.Print("Enter the Index of Left Node (-1 if NO) > ")
		fmt.Scan(&leftNode)
		if leftNode >= 0 {
			nodes[i].Left = &nodes[leftNode]
		}
		fmt.Print("Enter the Index of Right Node (-1 if NO) > ")
		fmt.Scan(&rightNode)
		if rightNode >= 0 {
			nodes[i].Right = &nodes[rightNode]
		}
		// Populate the Nodes

	}
	return nodes
}

func main() {
	// 5
	// 1 -1 -1
	// 3 -1 -1
	// 2 0 1
	// 6 -1 -1
	// 4 2 3
	nodes := read
}
