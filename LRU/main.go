package main

import "fmt"

type Qnode struct {
	val, key   int
	prev, next *Qnode
}

func newNode(key, val int) *Qnode {
	return &Qnode{
		key:  key,
		val:  val,
		prev: nil,
		next: nil,
	}
}

type Queue struct {
	head, tail *Qnode
}

func (q *Queue) addFrontPage(key, val int) *Qnode {
	page := newNode(key, val)
	if q.head == nil && q.tail == nil {
		q.head, q.tail = page, page
	} else {
		q.head.prev = page
		page.next = q.head
		q.head = page
	}

	return page
}

func (q *Queue) moveFront(page *Qnode) *Qnode {
	if q.head == nil && q.tail == nil {
		return nil
	} else if page == q.tail {
		q.tail = q.tail.prev
		q.tail.next = nil
	} else {
		page.prev.next = page.next
		page.next.prev = page.prev
	}

	q.head.prev = page
	page.next = q.head
	q.head = page
	return page
}

func (q *Queue) removeRear() {
	if q.head == nil && q.tail == nil {
	} else if q.tail == q.head {
		q.head, q.tail = nil, nil
	} else {
		q.tail = q.tail.prev
		q.tail.next = nil
	}

}

type LRU struct {
	size, capacity int
	queue          *Queue
	mp             map[int]*Qnode
}

func main() {
}
