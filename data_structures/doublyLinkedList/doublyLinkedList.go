package doublylinkedlist

import "fmt"

type Node struct {
	Next *Node
	Prev *Node
	Value int
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	len int
}

func (l *DoublyLinkedList) Len() int {
	return l.len
}

func (l *DoublyLinkedList) IsEmpty() bool {
	return l.len == 0
}

func (l *DoublyLinkedList) Head() (int, bool) {
	if l.len == 0 {
		return 0, false
	}
	return l.head.Value, true
}

func (l *DoublyLinkedList) Tail() (int, bool) {
	if l.len == 0 {
		return 0, false
	}
	return l.tail.Value, true
}

func (l *DoublyLinkedList) Clear() {
	l.head = nil
	l.tail = nil
	l.len = 0
}

func (l *DoublyLinkedList) PushBack(val int) {
	newNode := &Node{Next: nil, Prev: nil, Value: val}

	if l.len == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.Next = newNode
		newNode.Prev = l.tail
		l.tail = newNode
	}
	l.len++
}

func (l *DoublyLinkedList) PushFront(val int) {
	newNode := &Node{Next: nil, Prev: nil, Value: val}

	if l.len == 0 {
		l.head = newNode
		l.tail = newNode
	} else {
		l.head.Prev = newNode
		newNode.Next = l.head
		l.head = newNode
	}
	l.len++
}

func (l *DoublyLinkedList) PushAt(index int, val int) bool {
	if index < 0 || index > l.len {
		return false
	}
	if index == 0 {
		l.PushFront(val)
		return true
	}
	if index == l.len {
		l.PushBack(val) 
		return true
	}

	curNode := l.head

	for i := 0; i < index; i++ {
		curNode = curNode.Next
	}

	prevNode := curNode.Prev

	newNode := &Node{Next: curNode, Prev: prevNode, Value: val}
	prevNode.Next = newNode
	curNode.Prev = newNode
	
	return  true
}

func (l *DoublyLinkedList) RemoveAt(index int) bool {
	if l.len == 0 || index < 0 || index > (l.len-1) {
		return false
	}
	if index == 0 {
		_, _ = l.PopHead()
		return true
	}
	if index == l.len-1 {
		_, _ = l.PopTail()
		return true
	}

	curNode := l.head
	for i := 0; i < index-1; i++ {
		curNode = curNode.Next
	}

	nextNode := curNode.Next.Next
	curNode.Next = curNode.Next.Next
	nextNode.Prev = curNode

	return true
}

func (l *DoublyLinkedList) PopTail() (int, bool) {
	if l.len == 0 {
		return 0, false
	}
	res := l.tail.Value

	if l.len == 1 {
		l.Clear()
		return res, true
	}

	l.tail = l.tail.Prev
	l.tail.Next = nil
	l.len--

	return res, true
}

func (l *DoublyLinkedList) PopHead() (int, bool) {
	if l.len == 0 {
		return 0, false
	}
	res := l.head.Value
	
	if l.len == 1 {
		l.Clear()
		return res, true
	}
	
	l.head = l.head.Next
	l.head.Prev = nil
	l.len--

	return res, true
}

func (l *DoublyLinkedList) PopAt(index int) (int, bool) {
	if l.len == 0 || index < 0 || index > (l.len-1) {
		return 0, false
	}
	if index == 0 {
		val, ok := l.Head()
		return val, ok
	}
	if index == l.len - 1 {
		val, ok := l.Tail()
		return val, ok
	}

	curNode := l.head
	for i := 0; i < index; i++ {
		curNode = curNode.Next
	}

	result := curNode.Value

	l.RemoveAt(index)
	
	return result, true
}

func (l *DoublyLinkedList) Find(val int) bool {
	if l.len == 0 {
		return false
	}

	curNode := l.head
	for curNode != nil {
		if curNode.Value == val {
			return true
		}
		curNode = curNode.Next
	}

	return false
}

func (l *DoublyLinkedList) TraverseFromHead() {
	if l.len == 0 {
		return
	}

	curNode := l.head
	for curNode != nil {
		if curNode == l.tail {
			fmt.Print(curNode.Value)
			break
		}
		fmt.Print(curNode.Value, " <-> ")
		curNode = curNode.Next
	}
	fmt.Print("\n")
}

func (l *DoublyLinkedList) TraverseFromTail() {
	if l.len == 0 {
		return
	}

	curNode := l.tail
	for curNode != nil {
		if curNode == l.head {
			fmt.Print(curNode.Value)
			break
		}
		fmt.Print(curNode.Value, " <-> ")
		curNode = curNode.Prev
	}
	fmt.Print("\n")
}