package stack
import (
	"sync"
)

type node struct {
	item interface{}//元素
	next *node//指向下一个
}

type nodestack struct {
	first node
	number int
	lock  *sync.Mutex
}

func NewNodeStack() nodestack {
	s := nodestack{}
	s.lock=&sync.Mutex{}
	return s
}

func (s *nodestack)Push(value interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	oldFirst := s.first//把现在的栈顶存起来
	s.first = node{item:value, next:&oldFirst}//新建一个元素
	s.number++
}

func (s *nodestack)Pop() (interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.number==0 {
		return nil
	}

	item := s.first.item//将栈顶的元素取出
	s.first = *s.first.next//将栈顶指向到next中
	s.number--
	return item
}

func (s *nodestack)Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.number
}

func (s *nodestack)IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.number==0
}
