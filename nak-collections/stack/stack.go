package stack

import "fmt"

// A Node type to hold "any" value
// Can see prev and next node
type Node[T any] struct {
	prev  *Node[T]
	next  *Node[T]
	value *T
}

// An interface representing the Stack type
type IStack[T any] interface {
	push(newItem T)
	pop() T
	display()
	to_string() string
}

// A Stack type that can hold "any" one type of value
type Stack[T any] struct {
	head   *Node[T]
	length uint
}

// Create a new stack collection
func NewStack[T any](initValue T) *Stack[T] {
	head := new(Node[T])
	head.value = &initValue
	head.next = new(Node[T])
	return &Stack[T]{head: head, length: 1}
}

func (stack *Stack[T]) push(newItem T) {
	tempHead := stack.head
	stack.head = new(Node[T])
	stack.head.value = &newItem
	stack.head.next = tempHead
	stack.length++
}

func (stack *Stack[T]) pop() *T {
	to_pop := stack.head
	stack.head = stack.head.next
	stack.length--
	return to_pop.value
}

func (stack *Stack[T]) display() {
	tracker := stack.head
	for tracker.next != nil && !IsZero(tracker.value) {
		println(*tracker.value)
		tracker = tracker.next
	}
}

func (stack *Stack[T]) to_string() string {
	stack_display := "HEAD -> "
	tracker := stack.head
	for tracker.next != nil && !IsZero(tracker.value) {
		stack_display += fmt.Sprintf("%v -> ", *tracker.value)
		tracker = tracker.next
	}
	stack_display += "END"
	return stack_display
}

func Zero[T any]() T {
	return *new(T)
}

func IsZero[T comparable](v T) bool {
	return v == *new(T)
}
