package main

import (
	"fmt"
)

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	record := make(map[int]int)
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			record[v1+v2]++
		}
	}
	count := 0
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			if c, ok := record[0-v3-v4]; ok {
				count += c
			}
		}
	}
	return count
}

func main() {
	fmt.Println(fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
}
