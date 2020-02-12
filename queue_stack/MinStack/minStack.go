package MinStack

// No.155

type MinStack struct {
	stack []int
	min int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	s:=MinStack{
		stack: []int{},
		min:   0,
	}
	return s
}


func (this *MinStack) Push(x int)  {
	if len(this.stack)==0 {
		this.min=x
	}
	this.stack=append(this.stack,x)
	if x<this.min {
		this.min=x
	}
}


func (this *MinStack) Pop()  {
	this.stack= this.stack[:len(this.stack)-1]
	if len(this.stack)>0 {
		newmin:=this.stack[0]
		for _,v:=range this.stack {
			if v<=newmin {
				newmin=v
			}
		}
		this.min=newmin
	}

}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	return this.min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
