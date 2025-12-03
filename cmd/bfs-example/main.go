package bfsexample

import (
	"fmt"
	"github.com/Quasar777/go-algo/searches/bfs"
)

// GRAPH:
//          B  
//         /
//        A - D
//       /
//  E - C - F 

func main() {
	adjacencyList := make(map[rune][]rune)
	adjacencyList['A'] = []rune{'B', 'C', 'D'}
	adjacencyList['B'] = []rune{'A'}
	adjacencyList['C'] = []rune{'A', 'E', 'F'}
	adjacencyList['D'] = []rune{'A'}
	adjacencyList['E'] = []rune{'C'}
	adjacencyList['F'] = []rune{'C'}

	start := 'A'

	fmt.Println(bfs.BFS(adjacencyList, start, 'E')) // true
	fmt.Println(bfs.BFS(adjacencyList, start, 'I')) // false
}