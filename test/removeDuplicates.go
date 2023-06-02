package main

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	i := 0
	j := 1
	for j < len(nums) {
		if nums[i] == nums[j] {
			j++
		} else {
			i++
			nums[i] = nums[j]
			j++
		}
	}
	return i + 1
}
