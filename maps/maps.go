package main

import "fmt"
import "golang.org/x/tour/wc"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	wc.Test(WordCount)	
}

func WordCount(s string) map[string]int {
	return map[string]int {"x": 1}
}