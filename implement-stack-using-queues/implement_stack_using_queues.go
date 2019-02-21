type MyStack struct {
	stack []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	initStack := new(MyStack)
	initStack.stack = make([]int, 0, 0)
	return *initStack
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	last := len(this.stack)
	result := this.stack[last-1]
	this.stack = this.stack[:last-1]
	return result
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.stack[len(this.stack)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.stack) == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
