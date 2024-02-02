package lists

import (
	"errors"
	"fmt"
)

var EmptyList = errors.New("empty list")

type LinkedList[T any] struct {
	head, tail *node[T]
	length     int
}

func NewLinkedList[T any](data ...T) *LinkedList[T] {
	list := LinkedList[T]{}
	list.PushTail(data...)
	return &list
}

func (list *LinkedList[T]) Head() (T, error) {
	if list.head == nil {
		return *new(T), EmptyList
	}
	return list.head.data, nil
}

func (list *LinkedList[T]) Tail() (T, error) {
	if list.tail == nil {
		return *new(T), EmptyList
	}
	return list.tail.data, nil
}

func (list *LinkedList[T]) Len() int {
	return list.length
}

func (list *LinkedList[T]) PushTail(data ...T) {
	for _, data := range data {
		if list.head == nil {
			list.head = newNode(data)
			list.tail = list.head
		} else {
			list.tail.next = newNode(data)
			list.tail.next.prev = list.tail
			list.tail = list.tail.next
		}
		list.length += 1
	}
}

func (list *LinkedList[T]) PopTail() (T, error) {
	if list.tail == nil {
		return *new(T), EmptyList
	}

	list.length -= 1
	node := list.tail
	if node.prev == nil {
		list.tail = nil
		list.head = nil
		return node.data, nil
	}

	node.prev.next = nil
	list.tail = node.prev
	return node.data, nil
}

func (list *LinkedList[T]) PushHead(data ...T) {
	for _, data := range data {
		if list.head == nil {
			list.head = newNode(data)
			list.tail = list.head
		} else {
			list.head.prev = newNode(data)
			list.head.prev.next = list.head
			list.head = list.head.prev
		}
		list.length += 1
	}
}

func (list *LinkedList[T]) PopHead() (T, error) {
	if list.head == nil {
		return *new(T), EmptyList
	}

	list.length -= 1
	node := list.head
	if node.next == nil {
		list.tail = nil
		list.head = nil
		return node.data, nil
	}

	node.next.prev = nil
	list.head = node.next
	return node.data, nil
}

func (list *LinkedList[T]) String() string {
	repr := "["
	for e := list.head; e != nil; e = e.next {
		repr += fmt.Sprintf("%v", e.data)
		if e.next != nil {
			repr += " "
		}
	}
	return repr + "]"
}

func (list *LinkedList[T]) Iter() func() (T, error) {
	e := list.head
	return func() (T, error) {
		if e.next == nil {
			return *new(T), EmptyList
		}
		e = e.next
		return e.data, nil
	}
}

type node[T any] struct {
	prev, next *node[T]
	data       T
}

func newNode[T any](data T) *node[T] {
	return &node[T]{data: data}
}
