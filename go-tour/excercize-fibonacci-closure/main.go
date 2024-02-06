package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	current, next := 0, 1
	return func() (result int) {
		result = current
		current = next
		next = next + result
		return
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
