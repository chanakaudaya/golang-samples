package main

import (
	"fmt"
	"math"
	"math/rand"
)

var a, b, c bool

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func Sqrt(x float64) float64 {
	z := 1.0
	for ((x > z*z) && (x - z*z) > 0.000000001) || ((x <= z*z) && (z*z - x) > 0.000000001)  {
	z -= (z*z -x)/(2*z)
	}
	return z
}

func main() {
	fmt.Println("My favourite number is ", rand.Intn(100))
	fmt.Println(math.Pi)
	fmt.Println(add(23, 45))
	fmt.Println(swap("hello", "world"))
	var i int
	fmt.Println(i, a, b, c)
	fmt.Println(Sqrt(1600))

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

}
