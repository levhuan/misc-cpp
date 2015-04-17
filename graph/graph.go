/*
 * Graph representation (of M nodes and N edges):
 * - dense (fully-meshed and small)graph can be represented by MxN matrix
 * - sparse graph - adjacency list
 */
package main

import (
	"fmt"
	"time"
	"math/rand"
)

const (
	MAX_NODES = 200
	MAX_EDGES = MAX_NODES * MAX_NODES
)

var (
)

type vertice_id_t	int
type edge_id_t		int

type Edge struct {
	head	*Vertex
	tai		*Vertex
}

type Vertex struct {
	in	[]*Edge
	out	[]*Edge
}

type GraphOps interface {
	add_vertice() vertice_id_t
	del_vertice(id vertice_id_t) int
}

type Graph struct {
	num_vertex 	int
	vertices	[]*Vertex
	num_edges 	int
	edges		[]*Edge
}

func New(num_nodes int, num_edges int) *Graph {
	return &Graph{num_nodes, make([]*Vertex, num_nodes),
			num_edges, make([]*Edge, num_edges)}
}

func main() {
	fmt.Println("Main program!")
	rand.Seed(time.Now().UTC().UnixNano())

	num_nodes := rand.Intn(MAX_NODES)
	num_edges := rand.Intn(MAX_EDGES)

	graph := New(num_nodes, num_edges)
	graph.vertices[rand.Intn(num_nodes)] = &Vertex{}
	for i, vertex := range graph.vertices {
		if vertex != nil {
			fmt.Println("index: ", i, " vertex: ", vertex)
		}
	}
}
