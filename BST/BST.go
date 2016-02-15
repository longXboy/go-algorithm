package main

import (
	"fmt"
)

func main() {

	bst := BST{}
	bst.Set(2, "s2")
	bst.Set(1, "s1")
	bst.Set(3, "s3")
	bst.Set(4, "s4")
	bst.Set(-2, "s-2")
	bst.Set(-3, "s-3")
	bst.Set(-1, "s-1")
	bst.Set(5, "s5")
	bst.Set(9, "s9")
	bst.Set(0, "s0")
	fmt.Println(bst.DeleteByKey(2).key)
	print(bst.root)

}

func print(node *Node) {
	if node == nil {
		return
	}
	print(node.left)
	fmt.Println(node.key, " ", node.value, "  ", node.size())
	print(node.right)
}

type BST struct {
	root *Node
}

type Node struct {
	key   int
	value interface{}
	left  *Node
	right *Node
	nsize int
}

func (node *Node) size() int {
	if node == nil {
		return 0
	} else {
		return node.nsize
	}
}

func (bst *BST) DeleteByKey(key int) *Node {
	x := bst.get(bst.root, key)
	oldnode := *x
	*x = *bst.min(oldnode.right)
	x.right = bst.deleteMin(oldnode.right)
	x.left = oldnode.left
	return &oldnode
}

func (bst *BST) DeleteMin() {
	bst.root = bst.deleteMin(bst.root)
}

func (bst *BST) min(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return bst.min(node.left)
}

func (bst *BST) deleteMin(node *Node) *Node {
	if node.left == nil {
		return node.right
	}
	node.left = bst.deleteMin(node.left)
	node.nsize = node.left.size() + node.right.size()
	return node
}

func (bst *BST) SelectByRank(rank int) interface{} {
	node := bst.selectbyrank(bst.root, rank)
	if node == nil {
		return nil
	} else {
		return node.value
	}
}

func (bst *BST) selectbyrank(node *Node, rank int) *Node {
	if node == nil {
		return nil
	}
	if rank == node.left.size() {
		return node
	}
	if rank < node.left.size() {
		return bst.selectbyrank(node.left, rank)
	}
	if rank > node.left.size() {
		return bst.selectbyrank(node.right, rank-node.left.size()-1)
	}
	return nil
}

func (bst *BST) Get(key int) interface{} {
	node := bst.get(bst.root, key)
	if node == nil {
		return nil
	} else {
		return node.value
	}
}
func (bst *BST) get(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if node.key == key {
		return node
	} else if node.key > key {
		return bst.get(node.left, key)
	} else if node.key < key {
		return bst.get(node.right, key)
	}
	return nil
}

func (bst *BST) Set(key int, value interface{}) {
	if bst.root == nil {
		bst.root = &Node{key: key, value: value, nsize: 1}
	}
	bst.put(bst.root, key, value)
}

func (bst *BST) put(node *Node, key int, value interface{}) {
	if node.key == key {
		node.value = value
		return
	}
	if node.key > key {
		if node.left == nil {
			node.left = &Node{key: key, value: value, nsize: 1}
		} else {
			bst.put(node.left, key, value)
		}
	}
	if node.key < key {
		if node.right == nil {
			node.right = &Node{key: key, value: value, nsize: 1}
		} else {
			bst.put(node.right, key, value)
		}
	}
	node.nsize = node.left.size() + node.right.size() + 1
}
