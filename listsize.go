package piscine

/*
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

func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")

	fmt.Println(ListSize(link))
}
*/
type NodeF struct {
	Data interface{}
	Next *NodeF
}

type Listed struct {
	Head *NodeF
	Tail *NodeF
}

func ListSize(l *List) int {
	count := 0
	// go through the list increasing count
	for l.Head != nil {
		count++
		l.Head = l.Head.Next
	}
	return count
}
