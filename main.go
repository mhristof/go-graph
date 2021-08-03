package graph

// Map Dependency map information.
type DependencyMap struct {
	Nodes   []string // Nodes The nodes of the map.
	edge    map[string][]string
	visited map[string]bool
	sorted  []string
}

// Add Adds a node if it doesnt exist.
func (m *DependencyMap) AddNode(node string) {
	for _, item := range m.Nodes {
		if item == node {
			return
		}
	}

	m.Nodes = append(m.Nodes, node)
}

// Edge Adds an edge to the dependency list.
func (m *DependencyMap) Edge(source, dest string) {
	if m.edge == nil {
		m.edge = make(map[string][]string)
	}

	for _, node := range m.edge[source] {
		if node == dest {
			return
		}
	}

	m.edge[source] = append(m.edge[source], dest)
}

// SortAll Topological sort that includes the nodes not mentioned from the
// `startingNode`.
func (m *DependencyMap) SortAll(startingNode string) []string {
	return m.sort(startingNode, true)
}

// Sort Topological sort starting from node `startingNode`. It does not include
// nodes that are not connected to the `startingNode`.
func (m *DependencyMap) Sort(startingNode string) []string {
	return m.sort(startingNode, false)
}

func (m *DependencyMap) sort(startingNode string, all bool) []string {
	m.visited = nil
	m.sorted = nil

	m.sortRec(startingNode)

	for all {
		if len(m.sorted) == len(m.Nodes) {
			break
		}

		m.sortRec(diff(m.sorted, m.Nodes)[0])
	}

	return m.sorted
}

func diff(this, that []string) []string {
	items := map[string]int{}
	itemsThat := map[string]int{}

	var delta []string

	for _, item := range this {
		items[item] = 1
	}

	for _, item := range that {
		itemsThat[item] = 1

		if _, ok := items[item]; !ok {
			delta = append(delta, item)
		}
	}

	return delta
}

func (m *DependencyMap) sortRec(node string) {
	if _, ok := m.visited[node]; ok {
		return
	}

	if m.visited == nil {
		m.visited = make(map[string]bool)
	}

	m.visited[node] = true

	if m.edge[node] == nil {
		m.sorted = append(m.sorted, node)

		return
	}

	for _, subnode := range m.edge[node] {
		m.sortRec(subnode)
	}

	m.sorted = append(m.sorted, node)
}
