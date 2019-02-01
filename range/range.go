package main

import (
	"strings"
	"fmt"
)

var pow = []int{1,2,4,8,16,32,64,128}

func main() {
	for i,v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	s := "my name is chanaka is fernando is WSO2"
	result := strings.Fields(s)
	fmt.Println(result)
}