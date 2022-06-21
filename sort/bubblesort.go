package sort

func BubbleSort(nums []int) []int {
	r := len(nums)
	for i := 0; i < r; i++ {
		for j := i + 1; j < r; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
