package main

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
	record := make(map[int]int)
	for _, n := range nums1 {
		record[n] = 1
	}
	for _, n := range nums2 {
		if _, ok := record[n]; ok {
			record[n] = 0
		}
	}
	ret := make([]int, 0)
	for k, v := range record {
		if v == 0 {
			ret = append(ret, k)
		}
	}
	return ret
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
