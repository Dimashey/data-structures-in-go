package main

import (
	"fmt"

	"github.com/Dimashey/data-structures-in-go/lists"
	"github.com/Dimashey/data-structures-in-go/queues"
	"github.com/Dimashey/data-structures-in-go/sets"
	"github.com/Dimashey/data-structures-in-go/stacks"
)

func Arrays() {
	a := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(a)
}

func Slices() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	a = append(a, 9)

	fmt.Println(a)
}

func Sets() {
	set := sets.NewSet()

	set.Add("1")
	set.Add("1")
	set.Add("1")
	set.Add("1")
	set.Add("2")
	set.Add("3")

	set.Delete("2")

	set.List()

	fmt.Println(set.Contains("1"))
}

func Queues() {
	queue := queues.Queue{}

	fmt.Println(queue.IsEmpty())

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println(queue.Length())

	queue.Dequeue()

	element, _ := queue.Peek()

	fmt.Println(element)
}

func PriorityQueue() {
	priorityQueue := queues.PriorityQueue{}

	fmt.Println(priorityQueue.IsEmpty())

	priorityQueue.Enqueue(1, true)
	priorityQueue.Enqueue(2, false)
	priorityQueue.Enqueue(3, true)

	fmt.Println(priorityQueue.Lenght())

	priorityQueue.Dequeue()

	element, _ := priorityQueue.Peek()

	fmt.Println(element)
}

func Stacks() {
	stack := stacks.Stack{}

	fmt.Println(stack.IsEmpty())

	stack.Push(1)
	stack.Push(2)

	fmt.Println(stack.Length())

	elem, _ := stack.Pop()
	fmt.Println(elem)

	peek, _ := stack.Peek()
	fmt.Println(peek)
}

func LinkedList() {
	linkedList := lists.LinkedList{}

	linkedList.Insert("4")
	linkedList.Insert("3")
	linkedList.Insert("2")
	linkedList.Insert("1")

	linkedList.DeleteFirst()
	linkedList.Delete("3")

	if element := linkedList.Search("2"); element != nil {
		fmt.Println(element)
	}

	linkedList.List()
}

func main() {
	println("________________Array__________________")
	Arrays()
	println("________________Slice__________________")
	Slices()
	println("________________Sets___________________")
	Sets()
	println("________________Queues_________________")
	Queues()
	println("________________PriorityQueue__________")
	PriorityQueue()
	println("________________Stack__________________")
	Stacks()
	println("________________LinkedList_____________")
	LinkedList()
}
