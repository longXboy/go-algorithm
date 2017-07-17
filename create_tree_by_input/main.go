package main

import (
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func (head *Node) prevPrint() {
	fmt.Println(head.value)
	if head.left != nil {
		head.left.prevPrint()
	}
	if head.right != nil {
		head.right.prevPrint()
	}
}

func (head *Node) middlePrint() {
	if head.left != nil {
		head.left.middlePrint()
	}
	fmt.Println(head.value)
	if head.right != nil {
		head.right.middlePrint()
	}
}

func (head *Node) breadthPrint() {
	if head == nil {
		return
	}
	queue := make([]*Node, 0)
	lines := make([]int, 0)

	queue = append(queue, head)
	lines = append(lines, 0)
	lastLine := 0
	for {
		if lastLine != lines[0] {
			fmt.Println(" ")
		}
		first := queue[0]
		fmt.Printf("%d ", first.value)

		lastLine = lines[0]
		if first.left != nil {
			queue = append(queue, first.left)
			lines = append(lines, lastLine+1)
		}
		if first.right != nil {
			queue = append(queue, first.right)
			lines = append(lines, lastLine+1)
		}
		if len(queue) == 1 {
			break
		}
		queue = queue[1:]
		lines = lines[1:]
	}
}

func main() {
	aprev := []int{1, 2, 4, 7, 3, 5, 6, 8}
	amiddle := []int{4, 7, 2, 1, 5, 3, 8, 6}
	head := create(aprev, amiddle)
	head.breadthPrint()
}

func create(prev []int, middle []int) *Node {
	if len(prev) != len(middle) {
		return nil
	}
	if len(prev) == 0 {
		return nil
	}

	head := &Node{
		value: prev[0],
	}
	for i := range middle {
		if middle[i] == prev[0] {
			if i != 0 {
				head.left = create(prev[1:1+i], middle[:i])
			}
			if len(middle)-i > 1 {
				head.right = create(prev[1+i:], middle[i+1:])
			}
		}
	}
	return head
}
