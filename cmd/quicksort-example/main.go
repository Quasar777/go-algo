package main

import (
	"fmt"
	"github.com/Quasar777/go-algo/sorts/quick"
)

func main() {
	nums := []int{5, 3, 2, 14, 9, 0}

	quick.QuickSort(nums, 0, len(nums)-1)

	fmt.Println(nums)
}