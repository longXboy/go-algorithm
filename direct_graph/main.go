package main

import (
	"fmt"
)

type Digraph struct {
	V   int
	E   int
	adj [][]int
}

func (dg *Digraph) Init(v int) {
	dg.V = v
	dg.E = 0
	dg.adj = make([][]int, v)
}

func (dg *Digraph) addEdge(v int, w int) {
	for _, data := range dg.adj[v] {
		if data == w {
			return
		}
	}
	dg.adj[v] = append(dg.adj[v], w)
}

func (dg *Digraph) reverse() Digraph {
	newdg := Digraph{}
	newdg.Init(dg.V)
	for i := range dg.adj {
		for _, data := range dg.adj[i] {
			newdg.addEdge(data, i)
		}
	}
	return newdg
}

type DFS struct {
	marked []bool
}

func (d *DFS) Gen(g Digraph, v int) {
	d.marked = make([]bool, g.V)
	d.dfs(g, v)
}

func (d *DFS) dfs(g Digraph, v int) {
	d.marked[v] = true
	for _, data := range g.adj[v] {
		if !d.marked[data] {
			d.dfs(g, data)
		}
	}
}

type BFS struct {
	marked []bool
}

func (b *BFS) Gen(g Digraph, v int) {
	b.marked = make([]bool, g.V)
	var line []int
	b.marked[v] = true
	line = append(line, v)
	b.bfs(g, line)
}
func (b *BFS) bfs(g Digraph, line []int) {
	var newline []int
	for _, data := range line {
		for _, dataV := range g.adj[data] {
			if !b.marked[dataV] {
				b.marked[dataV] = true
				newline = append(newline, dataV)
			}
		}
	}
	if len(newline) > 0 {
		b.bfs(g, newline)
	}
}
func main() {
	graph := Digraph{}
	graph.Init(5)
	graph.addEdge(0, 1)
	graph.addEdge(0, 2)
	graph.addEdge(0, 3)
	graph.addEdge(4, 2)
	graph.addEdge(1, 4)
	bfs := BFS{}
	bfs.Gen(graph, 1)
	dfs := DFS{}
	dfs.Gen(graph, 1)
	fmt.Println(bfs.marked)
	fmt.Println(dfs.marked)
}
