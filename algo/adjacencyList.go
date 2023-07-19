package algo

import (
	"fmt"
	"strings"
)

type AdjacencyList map[string][]string
// Key: node name (string)
// Value: array of key's neighbors

func ReadAdjacencyList(input string, directed bool) AdjacencyList{
	adj := make(AdjacencyList)
	tokens := strings.Split(input, "\n")
	for _, token := range tokens {
		nodeNames := strings.Split(token, " ")
		first := string(nodeNames[0])
		second := string(nodeNames[1])

		if !contains(adj[first], second) {
			adj[first] = append(adj[first], second)
		}
		if !directed && !contains(adj[second], first) {
			adj[second] = append(adj[second], first)
		}
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

func contains(s []string, text string) bool {
	for _, element := range s {
		if element == text {
			return true
		} 
	}
	return false
}

