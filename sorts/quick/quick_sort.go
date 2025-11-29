package quick

func QuickSort(nums []int, start int, end int) {
	if start > end { return }

	left := start
	right := end
	pivot := nums[(start + end) / 2]

	for left < right {
		for nums[left] < pivot {
			left++
		}
		for nums[right] > pivot {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	if start < end {
		QuickSort(nums, start, right)
		QuickSort(nums, left, end)
	}
}