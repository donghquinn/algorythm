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

// 갱신
func (link *LinkedList) InsertFirst(node *Node) {
	if link.Header == nil {
		link.Header = node
		link.Tail = node
	} else {
		node.Next = link.Header
		link.Header = node
	}

	link.Count++
}

func (link *LinkedList) InsertLast(node *Node) {
	if link.Tail == nil {
		link.Header = node
		link.Tail = node
	} else {
		node.Next = link.Header
		link.Header = node
	}

	link.Count++
}

func (link *LinkedList) InsertMiddle(node *Node, prevNode *Node) bool {
	if prevNode == nil {
		return false
	} else {
		node.Next = prevNode.Next
		prevNode.Next = node
		if link.Tail == prevNode {
			link.Tail = node
		}

		link.Count++
		return true
	}
}

// 삭제
func (link *LinkedList) RemoveFirst() {
	if link.Tail == nil {
		return
	} else if link.Tail == link.Header {
		link.Tail = nil
		link.Header = nil
	} else {
		link.Header = link.Header.Next
	}
	link.Count--
}

func (link *LinkedList) RemoveLast() {
	if link.Tail == nil {
		return
	}

	if link.Tail == link.Header {
		link.Tail = nil
		link.Header = nil
	} else {
		for n := link.Header; n != nil; {
			if n.Next == link.Tail {
				n.Next = nil
				link.Tail = n
			}
			n = n.Next
		}
	}
	link.Count--
}

func (link *LinkedList) RemoveNode(node *Node) bool {
	if node == nil {
		return false
	}

	isRemove := false

	if link.Tail == link.Header {
		link.Count--
		link.Tail = nil
		link.Header = nil
		isRemove = true
	} else if link.Header == node {
		link.Header = node.Next
		isRemove = true
	} else {
		for n := link.Header; n != nil; {
			if n.Next == node {
				n.Next = node.Next
				isRemove = true
				if link.Tail == node {
					link.Tail = n
				}
			}
			n = n.Next
		}
	}

	if isRemove {
		link.Count--
	}

	return isRemove
}
