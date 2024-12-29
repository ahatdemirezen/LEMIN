package bfs

func Bfs(graph map[string]map[string]bool, start string, end string) [][]string {
	var allPaths [][]string
	queue := [][]string{{start}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		node := path[len(path)-1]

		if node == end {
			allPaths = append(allPaths, path)
			continue
		}

		for neighbor := range graph[node] {
			if !isNeighborInPath(neighbor, path) {
				newPath := make([]string, len(path))
				copy(newPath, path)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
	return allPaths
}

func isNeighborInPath(neighbor string, path []string) bool {
	for _, node := range path {
		if node == neighbor {
			return true
		}
	}
	return false
}
