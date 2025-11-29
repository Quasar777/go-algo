package main

import (
	"fmt"
	"github.com/Quasar777/go-algo/sorts/bubble"
)

func main() {
	nums := []int{5, 3, 2, 14, 9, 0}

	bubble.BubbleSort(nums)

	fmt.Println(nums)
}