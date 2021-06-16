package main

import "fmt"

func main() {
	r := peakIndexInMountainArray([]int{24,69,100,99,79,78,67,36,26,19})
	fmt.Println(r)
}

//二分
func peakIndexInMountainArray(arr []int) int {
	left, right, ans := 0, len(arr)-2, 0
	for left <= right{
		mid := (left + right) / 2
		if arr[mid] > arr[mid+1]{
			ans = mid
			right = mid - 1
		}else{
			left = mid + 1
		}
	}
	return ans
}

//枚举
func peakIndexInMountainArray_loop(arr []int) int {
	for i,n:=1,len(arr); i<n-1&&n>=3; i++{
		if arr[i] > arr[i-1] && arr[i] > arr[i+1]{
			return i
		}
	}
	return -1
}
