package queue

import (
	"container/list"
)

type ChanQueue struct {
	sem  chan bool
	list *list.List
}

type tFunc func(val interface{}) bool

func NewChanQueue() *ChanQueue {
	sem := make(chan bool, 1)
	list := list.New()
	return &ChanQueue{sem, list}
}

func (q *ChanQueue) Size() int {
	return q.list.Len()
}

func (q *ChanQueue) Enqueue(val interface{}) {
	q.sem <- true
	q.list.PushFront(val)
	<-q.sem
}

func (q *ChanQueue) Dequeue() interface{} {
	q.sem <- true
	e := q.list.Back()
	if e == nil {
		<-q.sem
		return nil
	}
	q.list.Remove(e)
	<-q.sem
	return e.Value
}

func (q *ChanQueue) Contain(val interface{}) bool {
	q.sem <- true
	e := q.list.Front()
	for e != nil {
		if e.Value == val {
			<-q.sem
			return true
		} else {
			e = e.Next()
		}
	}
	<-q.sem
	return false
}

func (q *ChanQueue) Query(queryFunc tFunc) *list.Element {
	q.sem <- true
	e := q.list.Front()
	for e != nil {
		if queryFunc(e.Value) {
			<-q.sem
			return e
		}
		e = e.Next()
	}
	<-q.sem
	return nil
}