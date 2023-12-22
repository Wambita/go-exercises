package piscine

/*
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

func main() {
	link := &List{}

	ListPushBack(link, "hello")
	ListPushBack(link, "how are")
	ListPushBack(link, "you")
	ListPushBack(link, 1)

	fmt.Println(ListAt(link.Head, 3).Data)
	fmt.Println(ListAt(link.Head, 1).Data)
	fmt.Println(ListAt(link.Head, 7))
}

// node
type NodeL struct {
	Data interface{}
	Next *NodeL
}

// list
type List struct {
	Head *NodeL
	Tail *NodeL
}
*/
func ListAt(l *NodeL, pos int) *NodeL {
	Current := l // start from head node
	curpos := 0  // track current position

	for Current != nil && curpos < pos { // loop to desired pos
		Current = Current.Next // next node
		curpos++               // increment pos

	}
	if Current == nil {
		return nil
	}
	return Current
}
