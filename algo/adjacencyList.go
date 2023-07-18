package algo

import (
	"fmt"
	"strings"
)

type AdjacencyList map[string][]string
// Key: node name (string)
// Value: array of key's neighbors

func ReadAdjacencyList(input string) AdjacencyList{
	adj := make(AdjacencyList)
	tokens := strings.Split(input, "\n")
	for _, token := range tokens {
		firstToken := string(token[0])
		secondToken := string(token[1])
		adj[firstToken] = append(adj[firstToken], secondToken)
	}
	return adj
}

func (adj AdjacencyList) Print() {
	println("==============")
	for key, value := range adj {
		fmt.Printf("%s: %s\n", key, value)
	}
	println("==============")
}