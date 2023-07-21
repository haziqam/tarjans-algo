package webapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/haziqam/tarjans-algo/packages/algo"
)

func handleFindBridge(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Content-Type", "text/plain")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body.", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var graphDataContainer [][]string
	json.Unmarshal(body, &graphDataContainer)
	fmt.Println("Graph data: ", graphDataContainer)

	startTime := time.Now()
	adjacencyList := algo.ReadAdjacencyList(graphDataContainer, false)
	fmt.Println("Adjacency list: ", adjacencyList)
	allNodes := algo.ReadAllNodes(graphDataContainer)
	solver := new(algo.BridgeSolver).Init(adjacencyList, allNodes)
	result := solver.GetBridge()
	executionTime := time.Since(startTime).Nanoseconds()
	fmt.Println(result)

	resultJSON := algoResult{
		Result: result,
		TimeNanoseconds: int(executionTime),
	}

	marshalledResult, err := json.Marshal(resultJSON)

	if (err != nil) {
		http.Error(w, "Error writing response.", http.StatusInternalServerError)
		return
	}

	w.Write(marshalledResult)
}