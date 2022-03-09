package main

import "fmt"

type Node struct {
	value interface{}
	left *Node
	right *Node
}


func (n *Node) InsertLeft(child *Node) {
	// MARK: -- if left is nil, add a node
	if n.left == nil {
		n.left = child
	} else {
		// MARK: -- original n.left will become the inserted child's left
		child.left = n.left
		// now n.left will equal the child
		n.left = child
	}
}

func (n *Node) InsertRight(child *Node) {
	if n.right == nil {
		n.right = child
	} else {
		child.right = n.right
		n.right = child
	}
}

func main() {
	root := Node{value: "a"}
	fmt.Println("root value", root.value)
	fmt.Println("root left before insert", root.left)

	root.InsertLeft(&Node{value: "b"})
	fmt.Println("root left after insert", root.left)

	root.InsertRight(&Node{value: "c"})
	fmt.Println("root right after insert", root.right)
	fmt.Println("root after insert", root.value, root.left, root.right)
	root.right.value = "hello"
	fmt.Println("root left before insert", root.right)

	tree := {
		"a", // Root
		{
			"b", // Left subtree
			{"d", {}, {}},
			{"e", {}, {}},
		},
		{
			"c", // Right subtree
			{"f", {}, {}},
			{} 
			
		}
	}
}