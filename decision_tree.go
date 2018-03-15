package main

import "fmt"

type Node struct {
	Left  *Node
	Value string
	Right *Node
}

var decisions = []bool{true, false, true, true}
var tree = Node{
	Value: "Did handins",
	Left:  nil,
	Right: &Node{
		Value: "Attended lectures",
		Right: nil,
		Left: &Node{
			Value: "Read Textbook",
			Left:  nil,
			Right: &Node{
				Value: "Made exercises",
				Left:  nil,
				Right: nil,
			},
		},
	},
}

func main() {
	// calling the walk method with index = 0
	Walk(&tree, 0)
}

// The method that 'walks' down the tree based on the decisions by the student
func Walk(n *Node, index int) {
	direction := decisions[index] // getting the boolean value from desicions starting from index 0

	if n.Left == nil && !direction { // if the Node below to the left is nil and we go in that direction we stop
		fmt.Println("Should not pass the exam")
		return
	}
	if n.Right == nil && direction { // if the Node below to the right is nil and we go in that direction we stop
		fmt.Println("Should pass the exam")
		return
	}
	index++ // incrementing the index parameter so that we get the next direction when we run Walk() again
	if !direction {
		fmt.Println("Went left")
		// if direction is false we go left in the Tree
		Walk(n.Left, index)
	} else if direction {
		// if direction is true we go right in the Tree
		fmt.Println("Went right")
		Walk(n.Right, index)
	}
}
