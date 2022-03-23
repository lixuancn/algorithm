package main

func xorOperation(n int, start int) int {
	ret := 0
	for i := 0; i < n; i++ {
		ret ^= start + 2*i
	}
	return ret
}
