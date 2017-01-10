package kit

import "errors"

type treeNode struct {
	left, right *treeNode
	val         int
}

//PriorityQueue is implemented with a BST tree
type PriorityQueue struct {
	root *treeNode
}

//Insert value into queue
func (q *PriorityQueue) Insert(value int) {
	currentNode := q.root
	if currentNode == nil {
		q.root = &treeNode{val: value}
		return
	}
	for {
		if currentNode.val < value {
			if currentNode.right == nil {
				currentNode.right = &treeNode{val: value}
				return
			}
			currentNode = currentNode.right
			continue
		}
		if currentNode.val > value {
			if currentNode.left == nil {
				currentNode.left = &treeNode{val: value}
				return
			}
			currentNode = currentNode.left
			continue
		}
	}
}

//GetMax value from the queue
func (q *PriorityQueue) GetMax() (int, error) {
	if q.root == nil {
		return 0, errors.New("Queue is empty")
	}
	current := q.root
	for current.right != nil {
		current = current.right
	}
	return current.val, nil
}

func (q *PriorityQueue) PopMax() (int, error) {
	if q.root == nil {
		return 0, errors.New("Queue is empty")
	}
	current := q.root
	for current.right != nil {
		current = current.right
	}
}
