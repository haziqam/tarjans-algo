package algo

import (
	"github.com/gammazero/deque"
)

// The SCC and the bridge solver is designed using OOP-like logic to
// avoid use of global variable and long parameter list.

type SccSolver struct {
	graph     	AdjacencyList
	allNodes 		[]string
	currentId 	int
	isOnStack 	map[string]bool
	ids       	map[string]int
	lowLink   	map[string]int
	sccCount 		int
	scc					[][]string
	stack 			deque.Deque[string]
	solved 			bool
}

const _UNVISITED int = -1

func (this *SccSolver) Init(graph AdjacencyList, allNodes []string) *SccSolver {
	this.graph = graph
	this.allNodes = allNodes
	this.solved = false
	return this
}

func (this *SccSolver) GetScc() ([][]string) {
	if !this.solved {this.solve()}
	return this.scc
}

func (this *SccSolver) solve() {
	this.ids = make(map[string]int)
	this.lowLink = make(map[string]int)
	this.isOnStack = make(map[string]bool)
	this.stack = deque.Deque[string]{}
	this.scc = make([][]string, 0)
	this.currentId = 0
	this.sccCount = 0

	for _, nodeName := range this.allNodes {
		this.isOnStack[nodeName] = false
		this.ids[nodeName] = _UNVISITED
	}

	for _, nodeName := range this.allNodes  {
		if this.ids[nodeName] == _UNVISITED {
			this.dfs(nodeName)
		}
	}

	this.solved = true
}

func (this *SccSolver) dfs(atNode string) {
	this.ids[atNode] = this.currentId
	this.lowLink[atNode] = this.currentId
	this.currentId++
	this.isOnStack[atNode] = true
	this.stack.PushBack(atNode)

	for _, neighbor := range this.graph[atNode] {
		if this.ids[neighbor] == _UNVISITED {
			this.dfs(neighbor)
		}
		if this.isOnStack[neighbor] {
			this.lowLink[atNode] = min(this.lowLink[atNode], this.lowLink[neighbor])
		}
	}

	if this.ids[atNode] == this.lowLink[atNode] {
		top := this.stack.PopBack();
		this.isOnStack[top] = false
		this.scc = append(this.scc, []string{})
		for {
			this.scc[this.sccCount] = append(this.scc[this.sccCount], top)
			if atNode == top {break}
			top = this.stack.PopBack() 
		}
		this.sccCount++
	}
}

func min(x int, y int) int {
	if x < y {return x}
	return y
}