# BFS - Breadth First Search

![bfs](../../docs/bfs.gif)

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
2. Use BFS function

```go
result := BFS(graph, 'a', 'g') // true
result := BFS(graph, 'a', 'k') // false
```