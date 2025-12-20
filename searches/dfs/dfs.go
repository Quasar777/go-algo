package dfs

func DFSIterative(graph map[rune][]rune, start rune, end rune) bool {
	visited := make(map[rune]bool) 
	stack := []rune{start}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == end {
			return true
		}

		if !visited[node] {
			visited[node] = true
			for _, neighbour := range graph[node] {
				if !visited[neighbour] {
					stack = append(stack, neighbour)
				}
			}
		}
	}

	return false
}

func DFSRecursive(graph map[rune][]rune, start rune, end rune, visited map[rune]bool) bool {
	if start == end {
		return true
	}

	if visited[start] {
		return false
	}

	visited[start] = true

	for _, neighbour := range graph[start] {
		if DFSRecursive(graph, neighbour, end, visited) {
			return true
		}
	}

	return false
}