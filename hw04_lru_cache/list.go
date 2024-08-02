package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	head, tail *ListItem
	size       int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.size
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	elem := &ListItem{Value: v}
	if l.head == nil {
		l.head = elem
		l.tail = elem
	} else {
		elem.Next = l.head
		l.head.Prev = elem
		l.head = elem
	}
	l.size++
	return elem
}

func (l *list) PushBack(v interface{}) *ListItem {
	elem := &ListItem{Value: v}
	if l.head == nil {
		l.head = elem
		l.tail = elem
	} else {
		elem.Prev = l.tail
		l.tail.Next = elem
		l.tail = elem
	}
	l.size++
	return elem
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.head = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.tail = i.Prev
	}
	l.size--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.head == i {
		return
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	if l.tail == i {
		l.tail = i.Prev
	}
	i.Prev = nil
	i.Next = l.head
	l.head.Prev = i
	l.head = i
}
