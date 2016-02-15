package main

import (
	"fmt"
)

func main() {
	rbt := RedBlackTree{}
	i := 0
	for {
		rbt.SetWithIteration(i, "t5")
		i++
		if i > 99 {
			break
		}
	}
	print(rbt.root)
	fmt.Println("-------")
	rbt.max(rbt.root)
}

func (rbt *RedBlackTree) max(node *Node) *Node {
	if node.right == nil {
		return node
	}
	fmt.Println(node.right)
	return rbt.max(node.right)
}

func (rbt *RedBlackTree) min(node *Node) *Node {
	if node.left == nil {
		return node
	}
	fmt.Println(node.left)
	return rbt.min(node.left)
}

func print(node *Node) {
	if node == nil {
		return
	}
	print(node.left)
	fmt.Println(node.key, " ", node.value, "  ", node.size())
	print(node.right)
}

type Node struct {
	key   int
	value interface{}
	left  *Node
	right *Node
	nsize int
	color bool //true:red false:black
}

func (node *Node) size() int {
	if node == nil {
		return 0
	} else {
		return node.nsize
	}
}

type RedBlackTree struct {
	root *Node
}
type NodeChain struct {
	node   *Node
	next   *NodeChain
	isLeft bool
}

func (rbt *RedBlackTree) SetWithIteration(key int, value interface{}) {
	if rbt.root == nil {
		rbt.root = &Node{key: key, value: value, nsize: 1, color: true}
		return
	}
	node := rbt.root
	var fisrt *NodeChain
	for {
		if node.key == key {
			node.value = value
			break
		} else if node.key > key {
			fisrt = &NodeChain{node: node, next: fisrt, isLeft: true}
			if node.left == nil {
				node.left = &Node{key: key, value: value, nsize: 1, color: true}
				break
			} else {
				node = node.left
			}

		} else if node.key < key {
			fisrt = &NodeChain{node: node, next: fisrt, isLeft: false}
			if node.right == nil {
				node.right = &Node{key: key, value: value, nsize: 1, color: true}
				break
			} else {
				node = node.right
			}

		}
	}
	for {
		if fisrt == nil {
			break
		}
		node = fisrt.node
		if node.right.Color() == true && node.left.Color() == false {
			if fisrt.next != nil {
				if fisrt.next.isLeft {
					fisrt.next.node.left = rbt.rotateLeft(node)
				} else {
					fisrt.next.node.right = rbt.rotateLeft(node)
				}
			} else {
				rbt.root = rbt.rotateLeft(node)
			}
		}
		if node.left.Color() == true && node.left.left.Color() == true {
			if fisrt.next != nil {
				if fisrt.next.isLeft {
					fisrt.next.node.left = rbt.rotateRight(node)
					fisrt.next.node.left.flipColors()
				} else {
					fisrt.next.node.right = rbt.rotateRight(node)
					fisrt.next.node.right.flipColors()
				}
			} else {
				rbt.root = rbt.rotateRight(node)
				rbt.root.flipColors()
			}
		}
		if node.left.Color() == true && node.right.Color() == true {
			node.flipColors()
		}
		node.nsize = node.left.size() + node.right.size() + 1
		fisrt = fisrt.next
	}
}

func (rbt *RedBlackTree) Set(key int, value interface{}) {
	rbt.root = rbt.put(rbt.root, key, value)
}

func (rbt *RedBlackTree) put(node *Node, key int, value interface{}) *Node {
	if node == nil {
		return &Node{key: key, value: value, nsize: 1, color: true}
	}
	if node.key == key {
		node.value = value
	}
	if node.key > key {
		node.left = rbt.put(node.left, key, value)
	} else if node.key < key {
		node.right = rbt.put(node.right, key, value)
	}
	if node.right.Color() == true && node.left.Color() == false {
		node = rbt.rotateLeft(node)
	}
	if node.left.Color() == true && node.left.left.Color() == true {
		node = rbt.rotateRight(node)
		node.flipColors()
	}
	if node.left.Color() == true && node.right.Color() == true {
		node.flipColors()
	}
	node.nsize = node.left.size() + node.right.size() + 1
	return node
}

func (node *Node) Color() bool {
	if node == nil {
		return false
	} else {
		return node.color
	}
}

func (node *Node) flipColors() {
	node.color = true
	node.left.color = false
	node.right.color = false
}

func (rbt *RedBlackTree) rotateLeft(h *Node) *Node {
	x := h.right
	h.right = x.left
	x.color, h.color = h.color, x.color
	x.left = h
	x.nsize = h.size()
	h.nsize = 1 + h.left.size() + h.right.size()
	return x
}

func (rbt *RedBlackTree) rotateRight(h *Node) *Node {
	x := h.left
	h.left = x.right
	x.color, h.color = h.color, x.color
	x.right = h
	x.nsize = h.size()
	h.nsize = 1 + h.left.size() + h.right.size()
	return x
}
