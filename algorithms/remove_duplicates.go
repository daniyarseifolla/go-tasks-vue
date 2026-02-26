package algorithms

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
