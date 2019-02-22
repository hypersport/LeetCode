type MyQueue struct {
	queue []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	initQueue := new(MyQueue)
	initQueue.queue = make([]int, 0)
	return *initQueue
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.queue = append(this.queue, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	tmp := this.queue[0]
	this.queue = this.queue[1:]
	return tmp
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.queue[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.queue) > 0 {
		return false
	}
	return true
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
