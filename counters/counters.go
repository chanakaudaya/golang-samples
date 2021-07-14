package main

import (
    "fmt"
    "sync/atomic"
)

var ops uint64

func main() {
	increment()
    fmt.Println("ops:", ops)
	increment()
    fmt.Println("ops:", ops)
	increment()
    fmt.Println("ops:", ops)
	increment()
    fmt.Println("ops:", ops)
}

func increment() uint64 {
	atomic.AddUint64(&ops, 1)
	return ops
}