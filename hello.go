package main

import (
	"fmt"
)

func hello() string {
	a := "hello Zhou"
	fmt.Printf("%s", a)
	return a
}

func main() {
	hello()
}
