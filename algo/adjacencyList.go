package algo

import (
	"fmt"
)

type AdjacencyList map[string][]string
// Key: node name (string)
// Value: array of key's neighbors

func ReadAdjacencyList(input [][]string, directed bool) AdjacencyList{
	adj := make(AdjacencyList)
	for _, pair := range input {
		first := pair[0]
		second := pair[1]

		if !contains(adj[first], second) {
			adj[first] = append(adj[first], second)
		}
		if !directed && !contains(adj[second], first) {
			adj[second] = append(adj[second], first)
		}
	}
	return adj
}

func ReadAllNodes(input [][]string) []string {
	allNodes := []string{}
	nodeSet := make(map[string]bool)

	for _, node := range input {
		first := node[0]
		second := node[1]

		if (!nodeSet[first]) {
			allNodes = append(allNodes, first)
			nodeSet[first] = true
		}
		if (!nodeSet[second]) {
			allNodes = append(allNodes, second)
			nodeSet[second] = true
		}
	}

	return allNodes
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

