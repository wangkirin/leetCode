package main

import "fmt"

func main() {

}

type MyStack struct {
	stack []int
}

func Constructor() MyStack {
	return MyStack{[]int{}}
}

func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyStack) Pop() int {
	if !this.Empty() {
		lens := len(this.stack)
		for i := 0; i < lens-1; i++ {
			x := this.stack[0]
			this.stack = this.stack[1:]
			this.stack = append(this.stack, x)
		}
		v := this.stack[0]
		this.stack = this.stack[1:]
		fmt.Println(this.stack)
		return v
	}
	return -1
}

func (this *MyStack) Top() int {
	if !this.Empty() {
		lens := len(this.stack)
		for i := 0; i < lens-1; i++ {
			x := this.stack[0]
			this.stack = this.stack[1:]
			this.stack = append(this.stack, x)
		}
		return this.stack[0]
	}
	return -1

}

func (this *MyStack) Empty() bool {
	if len(this.stack) == 0 {
		return true
	}
	return false
}
