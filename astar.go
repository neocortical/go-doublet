package main

import (
	"strings"
	"container/heap"
	"time"
	"fmt"
)

type Set map[string]*Node

func Doublet(a string, b string, dict *Dictionary, graph *Wordgraph) (*[]string, time.Duration) {
	startTime := time.Now()
	
	a = strings.ToLower(a)
	b = strings.ToLower(b)
	
	if len(a) != len(b) {
		fmt.Printf("Error: a and b must be the same length\n")
		return nil, time.Since(startTime)
	}
	if !dict.Contains(a) || !dict.Contains(b) {
		fmt.Printf("Error: a and b must be valid dictionary words\n")
		return nil, time.Since(startTime)
	}
	
	queue := &PriorityQueue{}
	heap.Init(queue)
	openset := make(Set)
	closedset := make(Set)
	node := Node{b, 0, estimateCost(b, a), nil, 0}
	heap.Push(queue, &node)
	openset[node.Word] = &node
	
	for queue.Len() > 0 {
		current := heap.Pop(queue).(*Node)
		delete(openset, current.Word)

		if current.Word == a {
			return reconstructPath(current), time.Since(startTime)
		}
		
		closedset[current.Word] = current
				
		// exapnd neighbors
		for _, neighbor := range (*graph)[current.Word] {
			if closedset[neighbor] != nil {
				continue
			}
			
			gval := current.Gval + 1
			neighborNode := openset[neighbor]
			if neighborNode == nil {
				neighborNode := Node{neighbor, gval, estimateCost(neighbor, a), current, 0}
				heap.Push(queue, &neighborNode)
				openset[neighborNode.Word] = &neighborNode
			} else if gval < neighborNode.Gval {
				neighborNode.Gval = gval
				neighborNode.Prev = current
				heap.Fix(queue, neighborNode.Index)
			}
		}
	}

	// failure state
	return nil, time.Since(startTime)
}

func estimateCost(from string, goal string) int {
	result := 0
	for i, _ := range from {
		if from[i] != goal[i] {
			result++
		}
	}

  return result
}

func reconstructPath(node *Node) *[]string {
	result := []string{}
	for node != nil {
		result = append(result, node.Word)
		node = node.Prev
	}
	
	return &result
}