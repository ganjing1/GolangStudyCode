package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Item struct {
	node     string
	priority int
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
func dijkstra(graph map[string]map[string]int, start string, end string) int {
	distances := make(map[string]int)
	for node := range graph {
		distances[node] = math.MaxInt32
	}
	distances[start] = 0
	queue := make(PriorityQueue, 0)
	heap.Init(&queue)
	heap.Push(&queue, &Item{node: start, priority: 0})
	for queue.Len() > 0 {
		currentItem := heap.Pop(&queue).(*Item)
		currentNode := currentItem.node
		currentDistance := currentItem.priority
		if currentNode == end {
			return distances[end]
		}
		if currentDistance > distances[currentNode] {
			continue
		}
		for neighbor, weight := range graph[currentNode] {
			distance := currentDistance + weight
			if distance < distances[neighbor] {
				distances[neighbor] = distance
				heap.Push(&queue, &Item{node: neighbor, priority: distance})
			}
		}
	}
	return -1
}
func main() {
	myGraph := map[string]map[string]int{
		"A": {"B": 5, "C": 2},
		"B": {"D": 4, "E": 2},
		"C": {"B": 8, "F": 7},
		"D": {"E": 6},
		"E": {},
		"F": {"E": 1},
	}
	fmt.Println(dijkstra(myGraph, "A", "E"))
}
