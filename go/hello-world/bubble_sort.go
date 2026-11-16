package main

// BubbleSort 将 nums 中的整数按升序原地排序。
func BubbleSort(nums []int) {
	for end := len(nums) - 1; end > 0; end-- {
		swapped := false
		for i := 0; i < end; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}
