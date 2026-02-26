package algorithms

import "math/rand"

// https://leetcode.com/problems/shuffle-an-array/

type Solution struct {
	original []int
	nums     []int
}

func Constructor(nums []int) Solution {
	original := make([]int, len(nums))
	copy(original, nums)
	return Solution{original: original, nums: nums}
}

func (s *Solution) Reset() []int {
	copy(s.nums, s.original)
	return s.nums
}

func (s *Solution) Shuffle() []int {
	for i := len(s.nums) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		s.nums[i], s.nums[j] = s.nums[j], s.nums[i]
	}
	return s.nums
}
