package sort

func partition(nums []int, l, r int) int {
	index := l - 1
	p := nums[r]
	for i := l; i < r; i++ {
		if nums[i] <= p {
			index += 1
			nums[index], nums[i] = nums[i], nums[index]
		}
	}
	nums[index+1], nums[r] = nums[r], nums[index+1]
	return index + 1
}

func quickSort(nums []int, l, r int) {
	if l < r {
		p := partition(nums, l, r)
		quickSort(nums, l, p-1)
		quickSort(nums, p+1, r)
	}
}

func QuickSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	quickSort(nums, 0, len(nums)-1)
}
