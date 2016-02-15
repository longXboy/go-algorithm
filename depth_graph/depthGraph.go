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
	g.Paths(0)
	fmt.Println(g.edgeTo)
}

type Bag struct {
	key  int
	next *Bag
}

type Graph struct {
	V      int //顶点数
	E      int //边数
	Bags   []*Bag
	edgeTo []int
	line   []int
	marked []bool
}

func (graph *Graph) Paths(start int) {
	graph.edgeTo = make([]int, graph.V)
	for i := range graph.edgeTo {
		graph.edgeTo[i] = -1
	}
	graph.paths(start, start)
}

func (graph *Graph) paths(vStart int, from int) {
	if graph.edgeTo[vStart] != -1 {
		return
	} else {
		graph.edgeTo[vStart] = from
	}
	node := graph.Bags[vStart]
	for {
		graph.paths(node.key, vStart)
		if node.next == nil {
			return
		} else {
			node = node.next
		}
	}
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

func (graph *Graph) Search(vStart int, vEnd int) {
	graph.line = make([]int, 0)
	graph.marked = make([]bool, graph.V)
	if graph.search(vStart, vEnd) {
		fmt.Println("line:", graph.line)
	}
}

func (graph *Graph) search(vStart int, vEnd int) bool {
	if graph.marked[vStart] {
		return false
	} else {
		graph.marked[vStart] = true
	}
	if vStart == vEnd {
		graph.line = append(graph.line, vStart)
		return true
	}
	graph.line = append(graph.line, vStart)
	node := graph.Bags[vStart]
	for {
		if graph.search(node.key, vEnd) {
			return true
		}
		if node.next == nil {
			graph.line = graph.line[:len(graph.line)-1]
			return false
		} else {
			node = node.next
		}
	}
}
