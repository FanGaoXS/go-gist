package main

import "fmt"

func New(cap int) *CircularQueue {
	return &CircularQueue{
		data:  make([]int, cap),
		front: 0,
		end:   0,
		cap:   cap,
	}
}

type CircularQueue struct {
	data []int

	front int // 指向队列头部
	end   int // 指向队列尾部的下一位置（待插入元素的位置）

	cap int // front 和 end 不断自增，通过对 cap 取模确定具体位置
}

func (c *CircularQueue) Push(val int) error {
	if c.IsFull() {
		return fmt.Errorf("queue is fulled")
	}

	index := c.end % c.cap
	c.data[index] = val

	c.end++
	return nil
}

func (c *CircularQueue) Pop() (int, error) {
	if c.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}

	index := c.front % c.cap
	val := c.data[index]

	c.front++
	return val, nil
}

func (c *CircularQueue) Front() (int, error) {
	if c.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}

	index := c.front % c.cap

	return c.data[index], nil
}

func (c *CircularQueue) End() (int, error) {
	if c.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}

	index := (c.end - 1) % c.cap

	return c.data[index], nil
}

func (c *CircularQueue) IsEmpty() bool {
	return c.end == c.front
}

func (c *CircularQueue) IsFull() bool {
	return c.end-c.front == c.cap
}

func (c *CircularQueue) Size() int {
	return c.end - c.front
}
