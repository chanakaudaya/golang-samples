package main

import (
	"fmt"
)

type ErrNegativeSqrt float64;

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if (x < 0) {
		er := ErrNegativeSqrt(x)
		return 0.0, er;
	} 
	z := 1.0
	for ((x > z*z) && (x - z*z) > 0.000000001) || ((x <= z*z) && (z*z - x) > 0.000000001)  {
	z -= (z*z -x)/(2*z)
	}
	return z, nil;
	
}


func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}