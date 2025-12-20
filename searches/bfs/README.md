# BFS - Breadth First Search

![bfs](../../docs/bfs.gif)

## Using

1. Init graph as a adjacency map

```go
graph := map[rune][]rune{}
graph['a'] = []rune{'b', 'c'}
graph['b'] = []rune{'e'}
graph['c'] = []rune{'d', 'f'}
graph['d'] = []rune{'e'}
graph['e'] = []rune{'g'}
graph['f'] = []rune{'e'}
graph['g'] = []rune{}
```
---
2. Use BFS function

```go
result := BFS(graph, 'a', 'g') // true
result := BFS(graph, 'a', 'k') // false
```