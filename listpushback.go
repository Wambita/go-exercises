package piscine

/*
func main() {
	link := &List{}

	ListPushBack(link, "Hello")
	ListPushBack(link, "man")
	ListPushBack(link, "how are you")

	for link.Head != nil {
		fmt.Println(link.Head.Data)
		link.Head = link.Head.Next
	}
}
*/
type NodeLs struct {
	Data interface{}
	Next *NodeLs
}

type Lists struct {
	Head *NodeLs
	Tail *NodeLs
}

// push new elem at the end of list
func ListPushBack(l *Lists, data interface{}) {
	newNode := &NodeLs{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
