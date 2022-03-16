package main

import "fmt"

type T struct {
	a int
	b string
}

func main() {
	var t1 *T
	var t2 *T
	fmt.Println(t1 == t2)
}
