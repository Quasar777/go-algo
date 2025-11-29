package merge

func mergeLists(nums1, nums2 []int) []int {
	result := make([]int, 0, len(nums1) + len(nums2))

	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	} 

	result = append(result, nums1[i:]...)
	result = append(result, nums2[j:]...)

	return result
}

func MergeSort(nums []int) []int {
	a := nums[:len(nums)/2]
	b := nums[len(nums)/2:]

	if len(a) > 1 {
		a = MergeSort(a)
	}
	if len(b) > 1 {
		b = MergeSort(b)
	}

	return mergeLists(a, b)
}