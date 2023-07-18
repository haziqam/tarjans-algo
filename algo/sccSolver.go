package algo

import (
	"github.com/gammazero/deque"
)

// The SCC solver is designed using OOP-like logic to avoid use
// of global variable and long parameter list.

type SccSolver struct {
	nodeCount int
	graph     AdjacencyList
	currentId int
	sccCount 	int
	isOnStack   map[string]bool
	ids       map[string]int
	lowLink   map[string]int
	sccs      map[string]string
	stack 		deque.Deque[string]
	solved 		bool
}

func (this *SccSolver) Init(graph AdjacencyList) *SccSolver {
	this.graph = graph
	this.nodeCount = len(graph)
	return this
}

func (this *SccSolver) GetScc() map[string]string {
	if !this.solved {this.solve()}
	return this.sccs
}

func (this *SccSolver) solve() {
	this.ids = make(map[string]int)
	this.lowLink = make(map[string]int)
	this.sccs = make(map[string]string)
	this.isOnStack = make(map[string]bool)
	this.stack = deque.Deque[string]{}
	for nodeName := range(this.graph) {
		this.isOnStack[nodeName] = false
	}

}

func (this *SccSolver) dfs() {

}