package main

import (
	"fmt"

	. "github.com/haziqam/tarjans-algo/packages/algo"
)

func main() {
	input := "A B\nB C\nC A\nB D\nD E\nE F\nF E"
	adjList := ReadAdjacencyList(input, true)
	adjList.Print()
	sccSolver := new(SccSolver).Init(adjList)
	sccResult := sccSolver.GetScc()
	fmt.Println(sccResult)

	adjList2 := ReadAdjacencyList(input, false)
	adjList2.Print()
	bridgeSolver := new(BridgeSolver).Init(adjList2)
	bridgeResult := bridgeSolver.GetBridge()
	fmt.Println(bridgeResult)
}