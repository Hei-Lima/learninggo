package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Pop(s *[]*tree.Tree) *tree.Tree {
	slice := *s
	lastIdx := len(slice) - 1

	poppedValue := slice[lastIdx]
	*s = slice[:lastIdx]

	return poppedValue
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var stack []*tree.Tree
	cur := t

	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = Pop(&stack)
			ch <- cur.Value

			cur = cur.Right
		}
	}

	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if !ok1 || !ok2 {
			return true
		}

		if v1 != v2 {
			return false
		}
	}
}

func main() {
	v := Same(tree.New(1), tree.New(1))
	fmt.Printf("%v", v)
}
