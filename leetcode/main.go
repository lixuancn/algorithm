package main

import "fmt"

type T struct {
	a int
	b string
}

func main() {
	t1 := T{a: 1, b: "1"}
	t2 := T{a: 1, b: "2"}
	fmt.Println(t1 == t2)
	a := make(map[int]bool, 1)
	a[0] = false
	A(a)
	fmt.Println(a)
}

func A(a map[int]bool) {
	a[0] = true
}
