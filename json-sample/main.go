package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m := Message{"Alice","Hello",1212121212}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(b)

	c := []byte(`{"Name":"Bob", "Body":"Pickle"}`)
	var n Message
	err = json.Unmarshal(c, &n)
	if (err != nil) {
		panic(err.Error())
	}
	fmt.Println(n)

}