package main

type MyStack struct {
	a []int
}

// LIFO queue to stack
// func Constructor() MyStack {
// 	return MyStack{}
// }
//
// func (this *MyStack) Push(x int) {
// 	this.a = append(this.a, x)
// 	for i := 0; i < len(this.a)-1; i++ {
// 		this.a = append(this.a, this.a[0])
// 		this.a = this.a[1:]
// 	}
// }
//
// func (this *MyStack) Pop() int {
// 	val := this.a[0]
// 	this.a = this.a[1:]
// 	return val
// }
//
// func (this *MyStack) Top() int {
// 	if len(this.a) == 0 {
// 		return 0
// 	}
// 	return this.a[0]
// }
//
// func (this *MyStack) Empty() bool {
// 	return len(this.a) == 0
// }

// Putting out seperately for easy access of all

// FIFO stack to queue
// type MyQueue struct {
// 	a []int
// }
//
// func Constructor() MyQueue {
// 	return MyQueue{}
// }
//
// func (this *MyQueue) Push(x int) {
// 	this.a = append(this.a, x)
// }
//
// func (this *MyQueue) Pop() int {
// 	val := this.a[0]
// 	this.a = this.a[1:]
// 	return val
// }
//
// func (this *MyQueue) Peek() int {
// 	if len(this.a) == 0 {
// 		return 0
// 	}
// 	return this.a[0]
// }
//
// func (this *MyQueue) Empty() bool {
// 	return len(this.a) == 0
// }

// min stack
type MinStack struct {
	nums []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.mins) == 0 || val <= this.GetMin() {
		this.mins = append(this.mins, val)
	}
	this.nums = append(this.nums, val)
}

func (this *MinStack) Pop() {
	if this.Top() == this.GetMin() {
		this.mins = this.mins[:len(this.mins)-1]
	}
	this.nums = this.nums[:len(this.nums)-1]
}

func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}
