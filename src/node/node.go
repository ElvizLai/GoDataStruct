package node

type Node struct {
	Item interface{}
	Next *Node
}

type DNode struct {
	Item   interface{}
	Before *Node
	Next   *Node
}

