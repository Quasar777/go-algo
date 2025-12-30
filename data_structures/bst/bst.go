package bst

import "fmt"

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

type BST struct {
	Root *Node
}

func (t *BST) Insert(val int) bool {
	newNode := &Node{Value: val}

	if t.Root == nil {
		t.Root = newNode
		return true
	}

	curNode := t.Root

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
	if t.Root == nil {
		return false
	}

	curNode := t.Root

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

func PrintInOrder(root *Node) {
	if root == nil {
		return
	}

	PrintInOrder(root.Left)
	fmt.Println(root.Value)
	PrintInOrder(root.Right)	
}

func PrintPreOrder(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node.Value)
	PrintPreOrder(node.Left)
	PrintPreOrder(node.Right)
}

func PrintPostOrder(node *Node) {
	if node == nil {
		return
	}

	PrintPreOrder(node.Left)
	PrintPreOrder(node.Right)
	fmt.Println(node.Value)
}