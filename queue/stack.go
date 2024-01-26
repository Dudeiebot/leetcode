package main

type MyStack struct {
	a []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.a = append(this.a, x)
	for i := 0; i < len(this.a)-1; i++ {
		this.a = append(this.a, this.a[0])
		this.a = this.a[1:]
	}
}

func (this *MyStack) Pop() int {
	val := this.a[0]
	this.a = this.a[1:]
	return val
}

func (this *MyStack) Top() int {
	if len(this.a) == 0 {
		return 0
	}
	return this.a[0]
}

func (this *MyStack) Empty() bool {
	return len(this.a) == 0
}

// Putting out seperately for easy access of all
