package concrete

import (
	"github.com/gonum/graph"
	"math"
)

// A dense graph is a graph such that all IDs are in a contiguous block from 0 to
// TheNumberOfNodes-1. It uses an adjacency matrix and should be relatively fast for both access
// and writing.
//
// This graph implements the CrunchGraph, but since it's naturally dense this is superfluous.
type DenseGraph struct {
	adjacencyMatrix []float64
	numNodes        int
}

// Creates a dense graph with the proper number of nodes. If passable is true all nodes will have
// an edge with cost 1.0, otherwise every node will start unconnected (cost of +Inf.)
func NewDenseGraph(numNodes int, passable bool) *DenseGraph {
	g := &DenseGraph{adjacencyMatrix: make([]float64, numNodes*numNodes), numNodes: numNodes}
	if passable {
		for i := range g.adjacencyMatrix {
			g.adjacencyMatrix[i] = 1.0
		}
	} else {
		for i := range g.adjacencyMatrix {
			g.adjacencyMatrix[i] = math.Inf(1)
		}
	}

	return g
}

func (g *DenseGraph) NodeExists(node graph.Node) bool {
	return node.ID() < g.numNodes
}

func (g *DenseGraph) Degree(node graph.Node) int {
	deg := 0
	for i := 0; i < g.numNodes; i++ {
		if g.adjacencyMatrix[i*g.numNodes+node.ID()] != math.Inf(1) {
			deg++
		}

		if g.adjacencyMatrix[node.ID()*g.numNodes+i] != math.Inf(1) {
			deg++
		}
	}

	return deg
}

func (g *DenseGraph) NodeList() []graph.Node {
	nodes := make([]graph.Node, g.numNodes)
	for i := 0; i < g.numNodes; i++ {
		nodes[i] = Node(i)
	}

	return nodes
}

func (g *DenseGraph) DirectedEdgeList() []graph.Edge {
	edges := make([]graph.Edge, 0, len(g.adjacencyMatrix))
	for i := 0; i < g.numNodes; i++ {
		for j := 0; j < g.numNodes; j++ {
			if g.adjacencyMatrix[i*g.numNodes+j] != math.Inf(1) {
				edges = append(edges, Edge{Node(i), Node(j)})
			}
		}
	}

	return edges
}

func (g *DenseGraph) Neighbors(node graph.Node) []graph.Node {
	neighbors := make([]graph.Node, 0)
	for i := 0; i < g.numNodes; i++ {
		if g.adjacencyMatrix[i*g.numNodes+node.ID()] != math.Inf(1) ||
			g.adjacencyMatrix[node.ID()*g.numNodes+i] != math.Inf(1) {
			neighbors = append(neighbors, Node(i))
		}
	}

	return neighbors
}

func (g *DenseGraph) EdgeBetween(node, neighbor graph.Node) graph.Edge {
	if g.adjacencyMatrix[neighbor.ID()*g.numNodes+node.ID()] != math.Inf(1) ||
		g.adjacencyMatrix[node.ID()*g.numNodes+neighbor.ID()] != math.Inf(1) {
		return Edge{node, neighbor}
	}

	return nil
}

func (g *DenseGraph) Successors(node graph.Node) []graph.Node {
	neighbors := make([]graph.Node, 0)
	for i := 0; i < g.numNodes; i++ {
		if g.adjacencyMatrix[node.ID()*g.numNodes+i] != math.Inf(1) {
			neighbors = append(neighbors, Node(i))
		}
	}

	return neighbors
}

func (g *DenseGraph) EdgeTo(node, succ graph.Node) graph.Edge {
	if g.adjacencyMatrix[node.ID()*g.numNodes+succ.ID()] != math.Inf(1) {
		return Edge{node, succ}
	}

	return nil
}

func (g *DenseGraph) Predecessors(node graph.Node) []graph.Node {
	neighbors := make([]graph.Node, 0)
	for i := 0; i < g.numNodes; i++ {
		if g.adjacencyMatrix[i*g.numNodes+node.ID()] != math.Inf(1) {
			neighbors = append(neighbors, Node(i))
		}
	}

	return neighbors
}

// DenseGraph is naturally dense, we don't need to do anything
func (g *DenseGraph) Crunch() {
}

func (g *DenseGraph) Cost(e graph.Edge) float64 {
	return g.adjacencyMatrix[e.Head().ID()*g.numNodes+e.Tail().ID()]
}

// Sets the cost of an edge. If the cost is +Inf, it will remove the edge,
// if directed is true, it will only remove the edge one way. If it's false it will change the cost
// of the edge from succ to node as well.
func (g *DenseGraph) SetEdgeCost(e graph.Edge, cost float64, directed bool) {
	g.adjacencyMatrix[e.Head().ID()*g.numNodes+e.Tail().ID()] = cost
	if !directed {
		g.adjacencyMatrix[e.Tail().ID()*g.numNodes+e.Head().ID()] = cost
	}
}

// Equivalent to SetEdgeCost(edge, math.Inf(1), directed)
func (g *DenseGraph) RemoveEdge(e graph.Edge, directed bool) {
	g.SetEdgeCost(e, math.Inf(1), directed)
}
