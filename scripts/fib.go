package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	arr := [2]int{0, 1}
	return func() int {
		atual := arr[0]
		arr[0], arr[1] = arr[1], arr[0]+arr[1]
		return atual
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
