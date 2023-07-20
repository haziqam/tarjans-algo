package main

import (
	"fmt"

	"github.com/haziqam/tarjans-algo/packages/webapi"
	// . "github.com/haziqam/tarjans-algo/packages/algo"
)

// func main() {
// 	input := [][]string{{"A", "B"}, {"C", "D"}}
// 	adjList := ReadAdjacencyList(input, true)
// 	nodes := ReadAllNodes(input)
// 	adjList.Print()
// 	fmt.Println("Nodes: \n", nodes)
// 	sccSolver := new(SccSolver).Init(adjList, nodes)
// 	sccResult := sccSolver.GetScc()
// 	fmt.Println(sccResult)

// 	adjList2 := ReadAdjacencyList(input, false)
// 	adjList2.Print()
// 	bridgeSolver := new(BridgeSolver).Init(adjList2, nodes)
// 	bridgeResult := bridgeSolver.GetBridge()
// 	fmt.Println(bridgeResult)
// }

func main() {
	err := webapi.StartServer("5000")
	if (err != nil) {
		fmt.Println(err)
	}
}