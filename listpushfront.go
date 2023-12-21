package piscine

/*
func main() {
	link := &Lists{}

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
type NodeS struct {
	Data interface{}
	Next *NodeS
}

// list structures
type Lists struct {
	Head *NodeS
	Tail *NodeS
}

func ListPushFront(l *List, data interface{}) {
	newNodes := &NodeS{Data: data, Next: nil}

	if l.Head == nil {
		l.Head = newNodes
		l.Tail = newNodes
	} else { // update Nextfield of newNode to point to the current head
		// update head to point to new node
		newNodes.Next = l.Head
		l.Head = newNodes

	}
}
