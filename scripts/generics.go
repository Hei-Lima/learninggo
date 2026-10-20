package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// ik that could've been a map
type Contact struct {
	name   string
	number int
}

func (c Contact) String() string {
	return fmt.Sprintf("%v: %v", c.name, c.number)
}

func readall[T any](l *List[T]) {
	for l != nil {
		fmt.Println(l.val)
		l = l.next
	}
}

func main() {
	root := List[Contact]{nil, Contact{"A", 0}}
	last := &root

	for i := 1; i < 5; i++ {
		cur := List[Contact]{nil, Contact{string('A' + i), i}}
		last.next = &cur
		last = &cur
	}

	readall(&root)
	return
}
