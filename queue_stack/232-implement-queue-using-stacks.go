package main

import (
	"container/list"
	"fmt"
)

func main() {

}

type MyQueue struct {
	len   int
	Queue *list.List
	Temp  *list.List
}

func ConstructorQueue() MyQueue {
	return MyQueue{
		len:   0,
		Queue: list.New(),
		Temp:  list.New(),
	}
}

func (this *MyQueue) Push(x int) {

	for this.Queue.Len() > 0 {
		t := this.Queue.Back()
		this.Queue.Remove(t)
		this.Temp.PushBack(t)
	}
	this.Queue.PushBack(x)
	for this.Temp.Len() > 0 {
		t := this.Temp.Back()
		this.Temp.Remove(t)
		this.Queue.PushBack(t)
	}
	this.len++

}

func (this *MyQueue) Pop() int {
	fmt.Println(this.Queue)
	v, _ := this.Queue.Back().Value.(int)
	this.Queue.Remove(this.Queue.Back())
	return v
}

func (this *MyQueue) Peek() int {
	fmt.Println(this.Queue)
	v, _ := this.Queue.Back().Value.(int)
	return v
}

func (this *MyQueue) Empty() bool {

	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
