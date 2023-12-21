package main

/*
func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "man")
	ListPushFront(link, "how are you")

	it := link.Head
	for it != nil {
		fmt.Print(it.Data, " ")
		it = it.Next
	}
	fmt.Println()
}
*/
// node structure
type NodeL struct {
	Data interface{}
	Next *NodeL
}

// list structures
type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data, Next: nil}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else { // update Nextfield of newNode to point to the current head
		// update head to point to new node
		newNode.Next = l.Head
		l.Head = newNode

	}
}
