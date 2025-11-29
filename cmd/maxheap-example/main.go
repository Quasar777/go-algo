package main

import (
	"fmt"
	"github.com/Quasar777/go-algo/data_structures/maxheap"
)

func main() {
	// Example of using a Maximum heap - Priority queue
	pQueue := &maxheap.MaxHeap{}

	pQueue.Insert(8)
	pQueue.Insert(2)
	pQueue.Insert(5)
	pQueue.Insert(12)
	pQueue.Insert(1)

	fmt.Println(pQueue.ExtractMax())
	fmt.Println(pQueue.ExtractMax())
	fmt.Println(pQueue.ExtractMax())
	fmt.Println(pQueue.ExtractMax())
	fmt.Println(pQueue.ExtractMax())
}