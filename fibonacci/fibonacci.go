package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(i int) int {
current := 1
previous := 0
return func(i int) int {
	if i < 2 {
		return i
	}
	temp := current
	current = previous + current
	previous = temp
	return current
}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
