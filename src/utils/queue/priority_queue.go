/**
 * Created by elvizlai on 2015/8/18 14:40
 * Copyright © PubCloud
 */
package queue
import (
	"container/heap"
)

type Item struct {
	Value    interface{}
	Priority int
}

type priorityQueue  []*Item

func (pq priorityQueue)Len() int {
	return len(pq)
}

func (pq priorityQueue)Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq priorityQueue)Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func NewPriorityQueue() *priorityQueue {
	pq := make(priorityQueue, 0)
	heap.Init(&pq)
	return &pq
}

func (pq *priorityQueue)Push(v interface{}) {
	item := v.(*Item)
	*pq = append(*pq, item)
}

//
func (pq *priorityQueue)Pop() interface{} {
	if pq.Len() == 0 {
		return nil
	}
	old := *pq
	n := len(old)
	item := old[n - 1]
	*pq = old[0 : n - 1]
	return item
}