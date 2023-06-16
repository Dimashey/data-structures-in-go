package main

import (
	"fmt"

	"github.com/Dimashey/data-structures-in-go/sets"
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

func main() {
	println("________________Array__________________")
	Arrays()
	println("________________Slice__________________")
	Slices()
	println("________________Sets___________________")
	Sets()
}
