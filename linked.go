package linked

type LinkedList struct {
	Header *Node
	Tail   *Node
	Count  int
}

type Node struct {
	Next  *Node
	Value interface{}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{nil, nil, 0}
}

func NewNode(v interface{}) *Node {
	return &Node{nil, v}
}
