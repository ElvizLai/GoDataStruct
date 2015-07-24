package queue

type arrayQueue struct {
	item []interface{}
	size int
}

func NewArrayQueue(cap int) *arrayQueue {
	return &arrayQueue{item: make([]interface{},0,cap), size:0}
}

func (a *arrayQueue)Enqueue(value interface{}) {
	a.item=append(a.item, value)
	a.size++
}

func (a *arrayQueue)Dequeue() interface{} {
	if a.size==0 {
		return nil
	}
	n := a.item[0]
	a.item=a.item[1:]
	a.size--
	return n
}

func (a *arrayQueue)Size() int {
	return a.size
}

func (a *arrayQueue)IsEmpty() bool {
	return a.size==0
}