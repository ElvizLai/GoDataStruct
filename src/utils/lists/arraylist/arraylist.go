/**
 * Created by elvizlai on 2015/9/1 15:18
 * Copyright © PubCloud
 */
package arraylist
import (
	"fmt"
	"strings"
)

type List struct {
	elements []interface{}
	size     int
}

const (
	GROWTH_FACTOR = float32(2.0)  // growth by 100%
	SHRINK_FACTOR = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

// Instantiates a new empty list
func New() *List {
	return &List{}
}

// Appends a value at the end of the list
func (list *List) Add(elements ...interface{}) {
	list.growBy(len(elements))
	for _, element := range elements {
		list.elements[list.size] = element
		list.size += 1
	}
}

// Returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (list *List) Get(index int) (interface{}, bool) {

	if !list.withinRange(index) {
		return nil, false
	}

	return list.elements[index], true
}

// Removes one or more elements from the list with the supplied indices.
func (list *List) Remove(index int) {

	if !list.withinRange(index) {
		return
	}

	list.elements[index] = nil                                    // cleanup reference
	copy(list.elements[index:], list.elements[index+1:list.size]) // shift to the left by one (slow operation, need ways to optimize this)
	list.size -= 1

	list.shrink()
}

// Check if elements (one or more) are present in the set.
// All elements have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (list *List) Contains(elements ...interface{}) bool {

	for _, searchElement := range elements {
		found := false
		for _, element := range list.elements {
			if element == searchElement {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Returns all elements in the list.
func (list *List) Values() []interface{} {
	newElements := make([]interface{}, list.size, list.size)
	copy(newElements, list.elements[:list.size])
	return newElements
}

// Returns true if list does not contain any elements.
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
	list.elements = []interface{}{}
}

// Swaps values of two elements at the given indices.
func (list *List) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}

func (list *List) String() string {
	str := "ArrayList\n"
	values := []string{}
	for _, value := range list.elements[:list.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is withing bounds of the list
func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size && list.size != 0
}

func (list *List) resize(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

// Expand the array if necessary, i.e. capacity will be reached if we add n elements
func (list *List) growBy(n int) {
	// When capacity is reached, grow by a factor of GROWTH_FACTOR and add number of elements
	currentCapacity := cap(list.elements)
	if list.size+n >= currentCapacity {
		newCapacity := int(GROWTH_FACTOR * float32(currentCapacity+n))
		list.resize(newCapacity)
	}
}

// Shrink the array if necessary, i.e. when size is SHRINK_FACTOR percent of current capacity
func (list *List) shrink() {
	if SHRINK_FACTOR == 0.0 {
		return
	}
	// Shrink when size is at SHRINK_FACTOR * capacity
	currentCapacity := cap(list.elements)
	if list.size <= int(float32(currentCapacity)*SHRINK_FACTOR) {
		list.resize(list.size)
	}

}