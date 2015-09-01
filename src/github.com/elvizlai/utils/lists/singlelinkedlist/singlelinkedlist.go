/**
 * Created by elvizlai on 2015/9/1 15:55
 * Copyright © PubCloud
 */
package singlelinkedlist

type List struct {
	first *element
	last  *element
	size  int
}

type element struct {
	next  *element
	value interface{}
}


func New() *List {
	return new(List)
}

//adding element to tail
func (list *List) Append(value interface{}) {
	e := &element{value:value}
	if list.size == 0 {
		list.first = e
		list.last = list.first
	}else {
		list.last.next = e
		list.last = list.last.next
	}
	list.size++
}

//adding element to head
func (list *List) Prepand(value interface{}) {
	e := &element{value:value}
	e.next = list.first
	list.first = e
	if list.size == 0 {
		list.last = list.first
	}
	list.size++
}

//get the index_th element
func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index){
		return nil, false
	}

	iter := list.first
	for e := 0; e != index; e, iter = e + 1, iter.next {
	}

	return iter.value, true
}

//retrun true if the list has the specified element
func (list *List) Contains(value interface{}) bool {
	if list.size == 0 {
		return false
	}

	for iter := list.first; iter != nil; iter = iter.next {
		if iter.value == value {
			return true
		}
	}
	return false
}

//remove the index_th element
func (list *List) Remove(index int) (interface{}, bool) {
	if !list.withinRange(index){
		return nil, false
	}

	var before *element

	iter := list.first
	for e := 0; e != index; e, iter = e + 1, iter.next {
		before = iter
	}

	//头 或尾巴
	if iter == list.first {
		list.first = iter.next
	}else if iter == list.last {
		list.last = before
	}

	if before != nil {
		before.next = iter.next
	}

	list.size--

	return iter.value, false
}


// Returns all elements in the list.
func (list *List) Values() []interface{} {
	values := make([]interface{}, list.size, list.size)
	for e, element := 0, list.first; element != nil; e, element = e + 1, element.next {
		values[e] = element.value
	}
	return values
}

//returns true if list is empty
func (list *List) Empty() bool {
	return list.size == 0
}

// Returns number of elements within the list.
func (list *List) Size() int {
	return list.size
}

// Removes all elements from the list.
func (list *List) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

// Check that the index is withing bounds of the list
func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size && list.size != 0
}