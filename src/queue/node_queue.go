package queue
import "sync"

type node struct {
	item interface{}//元素
	next *node//指向下一个
}

type nodequeue struct {
	head *node
	tail *node
	number int
	lock  *sync.Mutex
}

func NewNodeQueue() *nodequeue {
	q := &nodequeue{}
	q.lock = &sync.Mutex{}
	return q
}

func (q *nodequeue)Enqueue(value interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	n := &node{item:value}
	if q.tail==nil {
		q.head=n
		q.tail=n
	}else {
		q.tail.next=n
		q.tail=n
	}
	q.number++
}

func (q *nodequeue)Dequeue() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.head == nil {
		return nil
	}
	n := q.head
	q.head=n.next
	if q.head==nil {
		q.tail=nil
	}
	q.number--
	return n.item
}

func (q *nodequeue)IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.number==0
}

func (q *nodequeue)Size() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.number
}