package lists

import "fmt"

type ListNode struct {
	Value string
	Next  *ListNode
}

type LinkedList struct {
	Head *ListNode
	Size int
}

func (l *LinkedList) Insert(elem string) {
	node := ListNode{
		Value: elem,
		Next:  l.Head,
	}

	l.Head = &node
	l.Size++
}

func (l *LinkedList) DeleteFirst() {
	l.Head = l.Head.Next
	l.Size--
}

func (l *LinkedList) List() {
	current := l.Head

	for current != nil {
		fmt.Printf("%+v\n", current)

		current = current.Next
	}
}

func (l *LinkedList) Search(elem string) *ListNode {
	current := l.Head

	for current != nil {
		if current.Value == elem {
			return current
		}

		current = current.Next
	}

	return nil
}

func (l *LinkedList) Delete(elem string) {
	previous := l.Head
	current := l.Head

	for current != nil {
		if current.Value == elem {
			previous.Next = current.Next
			l.Size--
			return
		}

		previous = current
		current = current.Next
	}
}
