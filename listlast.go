package piscine

/*
func main() {
	link := &List{}
	link2 := &List{}

	ListPushBack(link, "three")
	ListPushBack(link, 3)
	ListPushBack(link, "1")

	fmt.Println(ListLast(link))
	fmt.Println(ListLast(link2))
}

// push new elem at the end of list
func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
*/
// node
type NodeP struct {
	Data interface{}
	Next *NodeP
}

// list
type ListP struct {
	Head *NodeP
	Tail *NodeP
}

func ListLast(l *ListP) interface{} {
	if l.Tail == nil {
		return nil
	} else {
		return l.Tail.Data
	}
}
