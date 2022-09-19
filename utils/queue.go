package utils

import (
	"sync"
)

const chunkSize int = 1024

type chunk struct {
	items [chunkSize]interface{}
	first int
	last  int
	next  *chunk
}

type Queue struct {
	head  *chunk
	tail  *chunk
	count int
	sync.Mutex
}

func NewQueue() *Queue {

	ck := new(chunk)
	queue := &Queue{
		head:  ck,
		tail:  ck,
		count: 0,
	}
	return queue
}

func (q *Queue) Push(item interface{}) {

	q.Lock()
	defer q.Unlock()

	if nil == item {
		return
	}

	if q.tail.last >= chunkSize {
		q.tail.next = new(chunk)
		q.tail = q.tail.next
	}

	q.tail.items[q.tail.last] = item
	q.tail.last++
	q.count++

}

func (q *Queue) Pop() interface{} {

	q.Lock()
	defer q.Unlock()

	if q.count == 0 {
		return nil
	}

	item := q.head.items[q.head.first]
	q.head.first++
	q.count--

	if q.head.first >= q.head.last {

		if q.count == 0 {
			q.head.first = 0
			q.head.last = 0
			q.head.next = nil
		} else {
			q.head = q.head.next
		}

	}

	return item

}

func (q *Queue) Len() int {

	q.Lock()
	defer q.Unlock()
	ct := q.count
	return ct
}
