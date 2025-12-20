# DFS - Depth First Search

![bfs](../../docs/dfs.gif)

## Using

1. Init graph as a adjacency map

```go
graph := map[rune][]rune{
    'a': {'b', 'c'},
    'b': {'e'},
    'c': {'d', 'f'},
    'd': {'e'},
    'e': {'g'},
    'f': {'e'},
    'g': {},
}
```
---
2. Use DFS function

```go
// Iterative
fmt.Println(dfsIterative(graph, 'a', 'b')) // true
fmt.Println(dfsIterative(graph, 'a', 'k')) // false

// Rec
visited := make(map[rune]bool) 
fmt.Println(dfsRecursive(graph, 'a', 'b', visited)) // true
fmt.Println(dfsRecursive(graph, 'a', 'k', visited)) // false
```