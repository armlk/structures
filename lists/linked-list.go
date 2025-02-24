package lists

import (
	"fmt"
)

type node[T comparable] struct {
	next *node[T]
	val  T
}

type linkedList[T comparable] struct {
	head *node[T]
}

// Initializes a linked list using the elements in the passed array
func LinkedList[T comparable](a []T) *linkedList[T] {
	l := &linkedList[T]{}
	for i := 0; i < len(a); i++ {
		l.Push(a[i])
	}

	return l
}

func (l *linkedList[T]) Add(ind uint, e T) (err error) {
	if ind > l.Size() {
		err = ListIndexOutOfBoundsError{ind, "Add"}
		return
	}
	new := &node[T]{nil, e}

	iter := l.head
	i := uint(0)
	for ; i != ind; i++ {
		iter = iter.next
	}

	if i == 0 && l.Size() == 0 {
		l.head = new
		return
	}

	if i == 0 {
		el, er := l.GetNode(0)
		if er != nil {
			err = er
			return
		}

		l.head = new
		l.head.next = el
		return
	}

	el, er := l.GetNode(i - 1)
	if er != nil {
		err = er
		return
	}

	el.next = new
	new.next = iter
	return
}

func (l *linkedList[T]) Delete(ind uint) (err error) {
	if ind >= l.Size() {
		err = ListIndexOutOfBoundsError{ind, "Delete"}
		return
	}

	var el, prev *node[T]
	for i := range l.Size() {
		el, err = l.GetNode(i)

		switch {
		case err != nil:
			return

		case i == ind && i == 0:
			l.head = el.next
			return

		case i == ind:
			prev, err = l.GetNode(i - 1)

			if err != nil {
				return
			}

			prev.next = el.next
			return
		}
	}

	return
}

func (l *linkedList[T]) Push(e T) (s uint) {
	new := &node[T]{nil, e}

	if l.head == nil {
		l.head = new
		s = l.Size()
		return
	}

	iter := l.head
	for iter.next != nil {
		iter = iter.next
	}

	iter.next = new
	s = l.Size()
	return
}

func (l *linkedList[T]) Pop() (s uint, err error) {
	if l.Size() == 0 {
		err = ListSizeTooSmallError("Pop")
		return
	}

	if l.Size() == 1 {
		l.head = nil
		s = 0
		return
	}

	el, er := l.GetNode(l.Size() - 2)
	if er != nil {
		err = er
		return
	}

	el.next = nil
	s = l.Size()
	return
}

func (l *linkedList[T]) Get(ind uint) (e T, err error) {
	if ind >= l.Size() {
		err = ListIndexOutOfBoundsError{ind, "Get"}
		return
	}

	iter := l.head
	for i := uint(0); ; i, iter = i+1, iter.next {
		if i == ind {
			e = iter.val
			break
		}

		if iter.next == nil {
			break
		}
	}

	return
}

func (l *linkedList[T]) IndexOf(e T) (ind uint, err error) {
	for i := range l.Size() {
		el, er := l.Get(i)
		if er != nil {
			err = er
			return
		}

		if el == e {
			ind = i
			return
		}
	}

	err = ElementNotFoundError[T]{e, "IndexOf"}
	return
}

func (l *linkedList[T]) GetNode(ind uint) (e *node[T], err error) {
	if ind >= l.Size() {
		err = ListIndexOutOfBoundsError{ind, "GetNode"}
		return
	}

	iter := l.head
	for i := uint(0); ; i++ {
		if i == ind {
			e = iter
			break
		}
		if iter.next == nil {

			break
		}

		iter = iter.next
	}

	return
}

func (l *linkedList[T]) Size() (i uint) {
	iter := l.head

	if iter == nil {
		return
	}

	for i = 1; iter.next != nil; i++ {
		iter = iter.next
	}

	return
}

func (e ListIndexOutOfBoundsError) Error() string {
	return fmt.Sprintf("index %v is invalid: index is out of bounds (%v)", e.ind, e.from)
}

func (e ElementNotFoundError[T]) Error() string {
	return fmt.Sprintf("could not find element %v (%v)", e.el, e.from)
}

func (e ListSizeTooSmallError) Error() string {
	return fmt.Sprintf("could not delete element at last index: list size too small (%v)", string(e))
}

func (l linkedList[T]) String() (s string) {
	s = "["
	for i := range l.Size() {
		el, err := l.Get(i)
		if err != nil {
			return
		}

		s += fmt.Sprintf("%v", el)

		if i < l.Size()-1 {
			s += ", "
		}
	}

	s += "]"
	return
}
