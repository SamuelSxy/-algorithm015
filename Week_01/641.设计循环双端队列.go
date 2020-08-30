/*
 * @lc app=leetcode.cn id=641 lang=golang
 *
 * [641] 设计循环双端队列
 */

// @lc code=start
type MyCircularDeque struct {
	max   int
	list  []int
	lenth int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		max:   k,
		list:  make([]int, 0), //slice append 从初始化尾部添加
		lenth: 0,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if !this.IsFull() {
		// this.list = append([]int{value}, this.list...)
		this.list = append(this.list, -1)
		copy(this.list[1:], this.list[0:])
		this.list[0] = value

		this.lenth++
		return true
	}
	return false
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if !this.IsFull() {
		this.list = append(this.list, value)
		this.lenth++
		return true
	}
	return false
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if !this.IsEmpty() {
		this.list = this.list[1:]
		this.lenth--
		return true
	}
	return false

}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if !this.IsEmpty() {
		this.list = this.list[:this.lenth-1]
		this.lenth--
		return true

	}
	return false
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if !this.IsEmpty() {
		return this.list[0]
	}
	return -1
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if !this.IsEmpty() {
		return this.list[this.lenth-1]
	}
	return -1
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if this.lenth == 0 {
		return true
	}
	return false
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if this.lenth == this.max {
		return true
	}
	return false
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end

