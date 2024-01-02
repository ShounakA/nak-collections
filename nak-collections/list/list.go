package list

import (
	"errors"
	"fmt"
)

// A Node type to hold "any" value
// Can see prev and next node
type Node[T any] struct {
	prev  *Node[T]
	next  *Node[T]
	value *T
}

type IList[T any] interface {
	append(newItem T)
	get(index uint) (T, error)

	remove(index uint) error
	put(newItem T, index uint)
	display()
	to_string() string
}

type List[T any] struct {
	head   *Node[T]
	length uint
}

// Create a new list collection
func NewList[T any]() *List[T] {
	head := new(Node[T])
	head.next = new(Node[T])
	return &List[T]{head: head, length: 0}
}

func (list *List[T]) display() {
	tracker := list.head
	for tracker.next != nil && !IsZero(tracker.value) {
		println(*tracker.value)
		tracker = tracker.next
	}
}

func (list *List[T]) to_string() string {
	list_display := "HEAD -> "
	tracker := list.head
	for tracker.next != nil && !IsZero(tracker.value) {
		list_display += fmt.Sprintf("%v -> ", *tracker.value)
		tracker = tracker.next
	}
	list_display += "END"
	return list_display
}

func (list *List[T]) append(newItem T) {
	if list.length == 0 {
		list.head.value = &newItem
		list.length++
	} else {
		lastNode := list.head
		for lastNode.next != nil && !IsZero(lastNode.value) {
			lastNode = lastNode.next
		}
		newNode := new(Node[T])
		lastNode.value = &newItem
		lastNode.next = newNode
		list.length++
	}
}

func (list *List[T]) remove(index uint) error {
	if list.length == 0 {
		return errors.New("cannot remove from empty list")
	}
	if index > list.length-1 {
		return errors.New(fmt.Sprintf("cannot remove from index %d. List size is %d", index, list.length))
	}
	if index == 0 {
		list.head = list.head.next
		list.length--
	} else {
		beforeNodeToDel := list.head
		count := uint(0)
		for beforeNodeToDel.next != nil && count < index-1 {
			beforeNodeToDel = beforeNodeToDel.next
			count++
		}
		nodeToDel := beforeNodeToDel.next
		beforeNodeToDel.next = nodeToDel.next
		list.length--
	}
	return nil
}

func (list *List[T]) get(index uint) (T, error) {
	if list.length == 0 {
		return Zero[T](), errors.New("cannot get from empty list")
	}
	if index > list.length-1 {
		return Zero[T](), errors.New(fmt.Sprintf("cannot get from index %d. List size is %d", index, list.length))
	}
	if index == 0 {
		return *list.head.value, nil
	} else {
		nodeToDel := list.head
		count := uint(0)
		for nodeToDel.next != nil && count < index {
			nodeToDel = nodeToDel.next
			count++
		}
		return *nodeToDel.value, nil
	}
}

func (list *List[T]) put(newItem T, index uint) error {
	if list.length == 0 {
		return errors.New("cannot get from empty list")
	}
	if index > list.length-1 {
		return errors.New(fmt.Sprintf("cannot get from index %d. List size is %d", index, list.length))
	}
	nodeToUpdate := list.head
	count := uint(0)
	for nodeToUpdate.next != nil && count < index {
		nodeToUpdate = nodeToUpdate.next
		count++
	}
	restofList := nodeToUpdate.next
	newNode := new(Node[T])
	newNode.value = &newItem
	newNode.next = restofList
	nodeToUpdate.next = newNode
	list.length++
	return nil
}

func Zero[T any]() T {
	return *new(T)
}

func IsZero[T comparable](v T) bool {
	return v == *new(T)
}
