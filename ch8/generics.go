package main

import (
	"fmt"
	"strconv"
)

// ----- 8.1 -----
// Write a generic function that doubles the value of any integer or float that’s
// passed in to it. Define any needed generic interfaces.

type IntAndFloat interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64
}

func DoubleValue[T IntAndFloat](num T) T {
	return num * 2
}

// ----- 8.2 -----
// Define a generic interface called Printable that matches a type that implements
// fmt.Stringer and has an underlying type of int or float64.
// Define types that meet this interface.
// Write a function that takes in a Printable and prints its value to the screen using fmt.Println.

type Printable interface {
	// we need ~ here since int has no methods, see page 193
	~int | ~float64
	fmt.Stringer
}

type AbbyInt int

func (a AbbyInt) String() string {
	// return fmt.Sprint(int(a))
	return strconv.Itoa(int(a))
}

func Print[T Printable](t T) {
	fmt.Println(t)
}

// ----- 8.3 -----
// Write a generic singly linked list data type.
// Each element can hold a comparable value and has a pointer to the next element in the list. The methods to implement are:
// adds a new element to the end of the linked list
// Add(T)
// adds an element at the specified position in the linked list
// Insert(T, int)
// returns the position of the supplied value, -1 if it's not present
// Index (T) int

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (l *LinkedList[T]) Add(t T) {
	node := &Node[T]{
		Value: t,
	}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}

	l.Tail.Next = node
	l.Tail = node
}

func (l *LinkedList[T]) Insert(t T, index int) {
	node := &Node[T]{
		Value: t,
	}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}

	if index <= 0 {
		node.Next = l.Head
		l.Head = node
		return
	}

	curNode := l.Head
	for i := 0; i < index-1; i++ {
		if curNode.Next == nil {
			curNode.Next = node
			l.Tail = node
			return
		}
		curNode = curNode.Next
	}

	node.Next = curNode.Next
	curNode.Next = node

	if l.Tail == curNode {
		l.Tail = node
	}

	return
}

func (l *LinkedList[T]) Index(t T) int {
	i := 0
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		if curNode.Value == t {
			return i
		}
		i++
	}

	return -1
}

func main() {
	// ----- 8.1 -----
	fmt.Println("----- 8.1 -----")
	fmt.Println(DoubleValue(5))
	fmt.Println(DoubleValue(1.9))

	// ----- 8.2 -----
	fmt.Println("----- 8.2 -----")
	var a AbbyInt = 519
	Print(a)

	// ----- 8.3 -----
	fmt.Println("----- 8.3 -----")
	l := &LinkedList[int]{}
	l.Add(5)
	l.Add(10)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))

	l.Insert(100, 0)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	l.Insert(200, 1)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(200))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Insert(300, 10)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Add(400)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Insert(500, 6)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}
}
