package main

import (
	"fmt"
	"time"
)

const ALPHA = "abcdefghijklmnopqrstuvwxyz"

type Wordgraph map[string][]string

func initWordgraph(dict *Dictionary) *Wordgraph {
	fmt.Printf("Loading word graph...\t")
	startTime := time.Now()

	graph := make(Wordgraph)

	for word, _ := range *dict {
		neighbors := []string{}
		for _, replacement := range ALPHA {
			for j := 0; j < len(word); j++ {
				newWord := word[:j] + string(replacement) + word[j+1:]
				if newWord != word && dict.Contains(newWord) {
					neighbors = append(neighbors, newWord)
				}
			}
		}
		graph[word] = neighbors
	}

	fmt.Printf("Loaded word graph in %d ms\n", time.Since(startTime)/1000000)

	return &graph
}
