package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	ClearFrontBack()
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len   int
	front *ListItem
	back  *ListItem
}

func (list *list) Len() int {
	return list.len
}

func (list *list) Front() *ListItem {
	return list.front
}

func (list *list) Back() *ListItem {
	return list.back
}

func (list *list) PushFront(v interface{}) *ListItem {
	var result *ListItem

	if list.front != nil {
		result = &ListItem{Next: list.front, Value: v}
		list.front.Prev = result
	} else {
		result = &ListItem{Value: v}
	}

	list.front = result

	if list.back == nil {
		list.back = result
	}

	list.len++

	return result
}

func (list *list) PushBack(v interface{}) *ListItem {
	var result *ListItem

	if list.back != nil {
		result = &ListItem{Prev: list.back, Value: v}
		list.back.Next = result
	} else {
		result = &ListItem{Value: v}
	}

	list.back = result

	if list.front == nil {
		list.front = result
	}

	list.len++

	return result
}

func (list *list) MoveToFront(i *ListItem) {
	list.Remove(i)
	list.PushFront(i.Value)
}

func (list *list) Remove(i *ListItem) {
	switch {
	case i.Prev != nil && i.Next != nil:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	case i.Prev != nil && i.Next == nil:
		i.Prev.Next = nil
		list.back = i.Prev
	case i.Next != nil && i.Prev == nil:
		i.Next.Prev = nil
		list.front = i.Next
	}

	list.len--
}

func (list *list) ClearFrontBack() {
	list.front = nil
	list.back = nil
	list.len = 0
}

func NewList() List {
	return new(list)
}
