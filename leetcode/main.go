package main

import "fmt"

type T struct {
	a int
	b string
}

func main() {
	a := make(map[int]bool, 1)
	a[0] = false
	A(a)
	fmt.Println(a)
}

func A(a map[int]bool) {
	a[0] = true
}
