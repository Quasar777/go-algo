package bfs

func BFS(graph map[rune][]rune, start rune, end rune) bool {
	visited := make(map[rune]bool) 
	visited[start] = true
	queue := []rune{start}

	for head := 0; head < len(queue); head++ {
		node := queue[head]

		for _, neighbour := range graph[node] {
			if neighbour == end {
				return true
			}
			if !visited[neighbour] {
				visited[neighbour] = true
				queue = append(queue, neighbour)
			}
		}
	}

	return false
}