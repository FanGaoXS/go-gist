package main

// https://leetcode.cn/problems/design-circular-queue/solutions/1715310/by-ac_oier-y11p/

type MyCircularQueue struct {
	data []int

	front int // 队列头部
	end   int // 队列尾部的下一位置（待插入元素的位置）

	cap int // 两变量始终自增，通过与 k 取模来确定实际位置。
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		data:  make([]int, k),
		front: 0,
		end:   0,
		cap:   k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	this.data[this.end%this.cap] = value
	this.end++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.front++
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	index := this.front % this.cap
	return this.data[index]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	index := (this.end - 1) % this.cap
	return this.data[index]
}

// IsEmpty 当 front 和 end 相等，队列存入元素和取出元素的次数相同，此时队列为空；
func (this *MyCircularQueue) IsEmpty() bool {
	return this.front == this.end
}

// IsFull 即队列元素个数，当元素个数为 k 个时，队列已满；
func (this *MyCircularQueue) IsFull() bool {
	return this.end-this.front == this.cap
}

func main() {
	queue := New(3)
	queue.Push(1)           // [1]
	queue.Push(2)           // [1,2]
	queue.Push(3)           // [1,2,3]
	queue.Push(4)           // [1,2,3]
	println(queue.End())    // [3]
	println(queue.IsFull()) // true
	queue.Pop()             // [2,3]
	queue.Push(4)           // [2,3,4]
	println(queue.Front())  // [2]
}
