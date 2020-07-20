package main

type Qnode struct {
	key, val   int
	prev, next *Qnode
}

func createNode(key, val int) *Qnode {
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
	page := createNode(key, val)
	if q.head == nil && q.tail == nil {
		q.head, q.tail = page, page

	} else {
		page.next = q.head
		q.head.prev = page
		q.head = page
	}
	return q.head
}

func (q *Queue) moveFrontPage(page *Qnode) *Qnode {
	if page == q.head {
		return q.head
	} else if page == q.tail {
		q.tail = q.tail.prev
		q.tail.next = nil
	} else {
		page.prev.next = page.next
		page.next.prev = page.prev
	}

	page.next = q.head
	q.head.prev = page
	q.head = page
	return q.head
}

func (q *Queue) removeRear() *Qnode {

}
