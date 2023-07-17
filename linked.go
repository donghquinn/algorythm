package linked

import "reflect"

type LinkedList struct {
	Header *Node
	Tail   *Node
	Count  int
}

type Node struct {
	Next  *Node
	Value interface{}
}

// 생성
func NewLinkedList() *LinkedList {
	return &LinkedList{nil, nil, 0}
}

func NewNode(v interface{}) *Node {
	return &Node{nil, v}
}

// 조회
func (link *LinkedList) FindByValue(value interface{}) *Node {
	for n := link.Header; n != nil; {
		if reflect.DeepEqual(n.Value, value) {
			return n
		} else {
			n = n.Next
		}
	}

	return nil
}

func (link *LinkedList) IndexOf(i int) *Node {
	node := link.Header
	if node == nil {
		return nil
	}
	for n := 0; n < i; n++ {
		if node == nil {
			return nil
		} else {
			node = node.Next
		}
	}
	return node
}
