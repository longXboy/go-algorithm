package main

import (
	"fmt"
)

func main() {
	g := Graph{}
	g.Init(6)
	g.AddTwoEdge(0, 1)
	g.AddTwoEdge(2, 3)
	g.AddTwoEdge(0, 3)
	g.AddTwoEdge(2, 4)
	g.AddTwoEdge(4, 5)
	g.AddTwoEdge(0, 5)
	g.ToString()
	b := BreathSearch{}
	b.Paths(&g, 0)
	fmt.Println(b.edgeTo)
}

type BreathSearch struct {
	queueHead *Bag
	edgeTo    []int
}

func (b *BreathSearch) Paths(g *Graph, start int) {
	b.edgeTo = make([]int, g.V)
	for i := range b.edgeTo {
		b.edgeTo[i] = -1
	}
	b.queueHead = &Bag{key: start, next: nil}
	b.edgeTo[start] = start
	for b.queueHead != nil {
		node := g.Bags[b.queueHead.key]
		for node != nil {
			if b.edgeTo[node.key] == -1 {
				b.edgeTo[node.key] = b.queueHead.key
				queue := b.queueHead
				for queue.next != nil {
					queue = queue.next
				}
				queue.next = &Bag{key: node.key, next: nil}
			}
			node = node.next
		}
		b.queueHead = b.queueHead.next
	}
}

type Bag struct {
	key  int
	next *Bag
}

type Graph struct {
	V    int //顶点数
	E    int //边数
	Bags []*Bag
}

func (graph *Graph) ToString() {
	for i := range graph.Bags {
		bag := graph.Bags[i]
		str := ""
		for bag != nil {
			str += fmt.Sprintf("%d", bag.key) + "   "
			bag = bag.next
		}
		fmt.Printf("%d : %s\n", i, str)
	}
}

func (graph *Graph) Init(nTop int) {
	graph.Bags = make([]*Bag, nTop)
	graph.V = nTop
	graph.E = 0
}

func (graph *Graph) AddTwoEdge(key1 int, key2 int) {
	graph.AddEdge(key1, key2)
	graph.AddEdge(key2, key1)

}

func (graph *Graph) AddEdge(key1 int, key2 int) {
	if key1 >= len(graph.Bags) {
		return
	}
	bag := graph.Bags[key1]
	if bag == nil {
		graph.Bags[key1] = &Bag{key: key2, next: nil}
		graph.E++
		return
	}
	for {
		if bag.key == key2 {
			break
		}
		if bag.next == nil {
			bag.next = &Bag{key: key2, next: nil}
			graph.E++
			break
		}
		bag = bag.next
	}
}
