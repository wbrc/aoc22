package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {

	var (
		graph              = make(map[position]*node)
		startNode, endNode *node
		endNodes           []*node
	)

	x := 0
	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		for y, h := range s.Text() {
			n := &node{
				elevation: getEvelation(h),
				pos:       position{x, y},
			}
			graph[position{x, y}] = n
			if h == 'S' {
				endNode = n
				endNodes = append(endNodes, n)
			} else if h == 'E' {
				startNode = n
			} else if h == 'a' {
				endNodes = append(endNodes, n)
			}
		}
		x++
	}

	for _, currentNode := range graph {
		for _, possibleNeighbor := range currentNode.possibleNeighbors() {
			if neighbor, ok := graph[possibleNeighbor]; ok && currentNode.reaches(neighbor) {
				currentNode.neighbors = append(currentNode.neighbors, neighbor)
			}
		}
	}

	for _, node := range graph {
		node.distance = math.MaxInt
	}

	startNode.distance = 0

	queue := make(nodeQueue, 0, len(graph))
	for _, node := range graph {
		node.index = len(queue)
		queue = append(queue, node)
	}
	heap.Init(&queue)

	for queue.Len() > 0 {
		currentNode := heap.Pop(&queue).(*node)
		if currentNode.distance == math.MaxInt {
			break
		}
		for _, neighbor := range currentNode.neighbors {
			if neighbor.index == -1 {
				continue
			}

			if alt := currentNode.distance + 1; alt < neighbor.distance {
				neighbor.distance = alt
				heap.Fix(&queue, neighbor.index)
			}
		}
	}

	visitableEndNodes := make([]*node, 0, len(endNodes))
	for _, n := range endNodes {
		if n.distance != math.MaxInt {
			visitableEndNodes = append(visitableEndNodes, n)
		}
	}

	sort.Slice(visitableEndNodes, func(i, j int) bool {
		return visitableEndNodes[i].distance < visitableEndNodes[j].distance
	})

	fmt.Println(visitableEndNodes[0].distance)
	_ = endNode
}

func getEvelation(r rune) int {
	if r == 'S' {
		return getEvelation('a')
	} else if r == 'E' {
		return getEvelation('z')
	}

	return int(r - 'a')
}

type position struct {
	x, y int
}

type node struct {
	index     int
	distance  int
	elevation int
	pos       position
	neighbors []*node
}

func (n *node) reaches(other *node) bool {
	heightDifference := other.elevation - n.elevation
	return -1 <= heightDifference
}

func (n *node) possibleNeighbors() []position {
	var neighbors []position
	if n.pos.x != 0 {
		neighbors = append(neighbors, position{n.pos.x - 1, n.pos.y})
	}
	if n.pos.y != 0 {
		neighbors = append(neighbors, position{n.pos.x, n.pos.y - 1})
	}
	neighbors = append(neighbors, position{n.pos.x + 1, n.pos.y})
	neighbors = append(neighbors, position{n.pos.x, n.pos.y + 1})
	return neighbors
}

type nodeQueue []*node

func (nq nodeQueue) Len() int { return len(nq) }

func (nq nodeQueue) Less(i, j int) bool {
	return nq[i].distance < nq[j].distance
}

func (nq nodeQueue) Swap(i, j int) {
	nq[i], nq[j] = nq[j], nq[i]
	nq[i].index = i
	nq[j].index = j
}

func (nq *nodeQueue) Push(x any) {
	n := len(*nq)
	item := x.(*node)
	item.index = n
	*nq = append(*nq, item)
}

func (nq *nodeQueue) Pop() any {
	old := *nq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*nq = old[0 : n-1]
	return item
}
