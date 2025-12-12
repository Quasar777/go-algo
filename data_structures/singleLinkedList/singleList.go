package singlelinkedlist

type Node struct {
	Next *Node
	Value int
}

type SingleLinkedList struct {
	head *Node
	tail *Node
}

func (l *SingleLinkedList) Push(val int) {
	newNode := &Node{Value: val, Next: nil}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.Next = newNode
		l.tail = newNode
	}
}

// Removes first founded element
func (l *SingleLinkedList) Remove(val int) bool {
	if l.head == nil {
		return false
	}

	if l.head.Value == val {
		if l.head.Next != nil {
			l.head = l.head.Next
		} else {
			l.head = nil
			l.tail = nil
		}
		return true
	}

	curr := l.head
	for curr.Next != nil && curr.Next.Value != val {
		curr = curr.Next
	}
	if curr.Next == nil {
		return false
	}
	
	if curr.Next == l.tail {
		l.tail = curr
		l.tail.Next = nil
		return true	
	}

	curr.Next = curr.Next.Next
	return true
}

func (l *SingleLinkedList) Find(val int) bool {
	if l.head == nil {
		return false
	}

	curr := l.head

	for curr != nil {
		if curr.Value == val {
			return true
		}
		curr = curr.Next
	}

	return false
}