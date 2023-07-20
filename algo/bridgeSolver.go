package algo

type BridgeSolver struct {
	graph     AdjacencyList
	allNodes  []string
	currentId int
	isVisited map[string]bool
	ids       map[string]int
	lowLink   map[string]int
	bridges   [][]string
	solved    bool
}

const _NOPARENT string = "-"

func (this *BridgeSolver) Init(graph AdjacencyList, allNodes []string) *BridgeSolver {
	this.graph = graph
	this.allNodes = allNodes
	this.solved = false
	return this
}

func (this *BridgeSolver) GetBridge() [][]string {
	if !this.solved {
		this.solve()
	}
	return this.bridges
}

func (this *BridgeSolver) solve() {
	this.currentId = 0
	this.lowLink = make(map[string]int)
	this.ids = make(map[string]int)
	this.isVisited = make(map[string]bool)
	this.bridges = make([][]string, 0)

	for _, nodeName := range this.allNodes {
		this.isVisited[nodeName] = false
	}

	for _, nodeName := range this.allNodes {
		if !this.isVisited[nodeName] {
			this.dfs(nodeName, _NOPARENT)
		}
	}

	this.solved = true
}

func (this *BridgeSolver) dfs(atNode string, parent string) {
	this.ids[atNode] = this.currentId
	this.lowLink[atNode] = this.currentId
	this.currentId++
	this.isVisited[atNode] = true

	for _, neighbor := range this.graph[atNode] {
		if neighbor == parent {
			continue
		}
		if !this.isVisited[neighbor] {
			this.dfs(neighbor, atNode)
			this.lowLink[atNode] = min(this.lowLink[atNode], this.lowLink[neighbor])

			if this.ids[atNode] < this.lowLink[neighbor] {
				// If there's no back edge between atNode and neighbor
				// => Edge(atNode, neighbor) is a bridge
				this.bridges = append(this.bridges, []string{atNode, neighbor})
			}
		} else {
			this.lowLink[atNode] = min(this.lowLink[atNode], this.lowLink[neighbor])
		}
	}
}
