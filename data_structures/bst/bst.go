package bst

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

type BST struct {
	root *Node
}

func (t *BST) Insert(val int) bool {
	newNode := &Node{Value: val}

	if t.root == nil {
		t.root = newNode
		return true
	}

	curNode := t.root

	for {
		if newNode.Value > curNode.Value {
			if curNode.Right == nil {
				curNode.Right = newNode
				return true
			}
			curNode = curNode.Right
		} else if newNode.Value < curNode.Value {
			if curNode.Left == nil {
				curNode.Left = newNode
				return true
			}
			curNode = curNode.Left
		} else {
			return false
		}
	}
}

func (t *BST) Search(val int) bool {
	if t.root == nil {
		return false
	}

	curNode := t.root

	for curNode != nil {
		if curNode.Value == val {
			return true
		}
		if val > curNode.Value {
			curNode = curNode.Right
		} else {
			curNode = curNode.Left
		}
	}

	return false
}