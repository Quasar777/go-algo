package dijkstra

type Edge struct {
    To   rune
    Cost int
}

func dijkstra(graph map[rune][]Edge, start rune) map[rune]int {
	// Иницилизация
	distances := map[rune]int{}
	visited := map[rune]bool{}
	
	// Помечаем минимальне пути до всех вершин как infinity
	// Длина пути до стартовой вершины - 0
	for vertex := range graph {
		distances[vertex] = 10000000000
	}
	distances[start] = 0

	// Основной цикл
	for len(visited) < len(graph) {
		closestVertex := ' '
		smallestDist := 10000000000

		for vertex := range distances {
			if !visited[vertex] && distances[vertex] < smallestDist {
				smallestDist = distances[vertex]
				closestVertex = vertex
			}
		}

		if closestVertex == ' ' { 
			break 
		}

		visited[closestVertex] = true

		for _, neighbour := range graph[closestVertex] {
			weight := neighbour.Cost
			newWeight := distances[closestVertex] + weight
			if newWeight < distances[neighbour.To] {
				distances[neighbour.To] = newWeight
			}
		}
	}

	return distances
}