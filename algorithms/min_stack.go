package algorithms

// https://leetcode.com/problems/min-stack/

type MinStack struct {
	main []int
	min  []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.main = append(this.main, val)

	n := len(this.min)
	if n == 0 || val < this.min[n-1] {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, this.min[n-1])
	}
}

func (this *MinStack) Pop() {
	n := len(this.main)
	this.main = this.main[:n-1]
	this.min = this.min[:n-1]
}

func (this *MinStack) Top() int {
	return this.main[len(this.main)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
