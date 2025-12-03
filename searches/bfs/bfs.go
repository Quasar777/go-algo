package bfs

import "github.com/Quasar777/go-algo/data_structures/queue"

func BFS(graph map[rune][]rune, start rune, target rune) bool {
	visited := make(map[rune]bool)
	queue := &queue.Queue[rune]{Nums: []rune{}}

	queue.Enque(start)
	visited[start] = true

	current, ok := queue.Deque()
	if current == target { return true}
	for ok {
		for _, r := range graph[current] {
			_, isVisited := visited[r]
			if !isVisited {
				queue.Enque(r)
				visited[r] = true
			}
		}
		current, ok = queue.Deque()
		if current == target {return true}
	}

	return false
}