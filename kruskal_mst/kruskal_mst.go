package main

import (
	"fmt"
	"sort"
)

type Edge struct {
	a, b int
	w    int
}

type Graph struct {
	edges   []Edge
	n, m    int
	parents []int
	weights []int
}

func (g *Graph) Init(n, m int) {
	g.n, g.m = n, m
	g.parents = make([]int, n)
	g.weights = make([]int, n)
	for i := range g.parents {
		g.parents[i] = i
	}
}

func (g *Graph) Find(x int) int {
	if g.parents[x] != x {
		g.parents[x] = g.Find(g.parents[x])
	}
	return g.parents[x]
}

func (g *Graph) Union(x, y int) {
	x, y = g.Find(x), g.Find(y)
	if g.weights[x] > g.weights[y] {
		g.parents[y] = x
	} else {
		g.parents[x] = y
		if g.weights[x] == g.weights[y] {
			g.weights[y]++
		}
	}
}

func (g *Graph) AddEdge(a, b, w int) {
	g.edges = append(g.edges, Edge{a, b, w})
}

func (g *Graph) Kruskal() {
	sort.Slice(g.edges, func(i, j int) bool {
		return g.edges[i].w < g.edges[j].w
	})
	for i := range g.edges {
		if g.Find(g.edges[i].a) != g.Find(g.edges[i].b) {
			fmt.Println(g.edges[i].a, g.edges[i].b)
			g.Union(g.edges[i].a, g.edges[i].b)
		}
	}
}

func main() {
	g := &Graph{}
	g.Init(4, 5)
	g.AddEdge(0, 1, 10)
	g.AddEdge(0, 2, 6)
	g.AddEdge(0, 3, 5)
	g.AddEdge(1, 3, 15)
	g.AddEdge(2, 3, 4)
	g.Kruskal()
}
