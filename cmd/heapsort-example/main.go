package main

import (
	"fmt"
	"github.com/Quasar777/go-algo/sorts/heap"
)

func main() {
	nums := []int{5, 3, 2, 14, 9, 0}

	h := &heap.MaxHeap{}
	h.Build(nums)
	heap.HeapSort(h)
	copy(nums, h.Nums)

	fmt.Println(nums)
}