package maxheap

import "fmt"

type MaxHeap struct {
	Nums []int
}

func (h *MaxHeap) Build(nums []int) {
	h.Nums = make([]int, len(nums))
	copy(h.Nums, nums)
	
	size := len(h.Nums)

	lastNodeWithoutLeafIndex := (size / 2) - 1

	for i := lastNodeWithoutLeafIndex; i >= 0; i-- {
		h.HeapifyDown(i)
	}
}

func (h *MaxHeap) ParentIndex(index int) int {
	return (index - 1) / 2
}

func (h *MaxHeap) LeftChildIndex(index int) int {
	return index*2 + 1
}

func (h *MaxHeap) RightChildIndex(index int) int {
	return index*2 + 2
}

func (h *MaxHeap) Print() {
	fmt.Println(h.Nums)
}

func (h *MaxHeap) Insert(value int) {
	h.Nums = append(h.Nums, value)

	if len(h.Nums) < 2 {
		return
	}

	valueIndex := len(h.Nums) - 1
	parentIndex := h.ParentIndex(valueIndex)

	for valueIndex != 0 && h.Nums[parentIndex] < h.Nums[valueIndex] {
		h.Nums[parentIndex], h.Nums[valueIndex] = h.Nums[valueIndex], h.Nums[parentIndex]
		valueIndex = parentIndex
		parentIndex = h.ParentIndex(valueIndex)
	}
}

func (h *MaxHeap) HeapifyDown(index int) {
	largest := index
	size := len(h.Nums)

	for {
		left := h.LeftChildIndex(index)
		right := h.RightChildIndex(index)

		if left < size && h.Nums[left] > h.Nums[largest] {
			largest = left
		}
		if right < size && h.Nums[right] > h.Nums[largest] {
			largest = right
		}
		if largest != index {
			h.Nums[largest], h.Nums[index] = h.Nums[index], h.Nums[largest]
			index = largest
		} else {
			break
		}
	}
}

func (h *MaxHeap) ExtractMax() int {
	if len(h.Nums) == 0 { return 0 }

	max := h.Nums[0]
	h.Nums[0], h.Nums[len(h.Nums)-1] =  h.Nums[len(h.Nums)-1], h.Nums[0]
	h.Nums = h.Nums[:len(h.Nums)-1]
	h.HeapifyDown(0)

	return max
}