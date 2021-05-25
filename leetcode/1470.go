package main
import (
	"fmt"
)

func main() {
	r := shuffle([]int{2,5,1,3,4,7}, 3)
	fmt.Println(r)
}

//快慢指针
func shuffle(nums []int, n int) []int {
	var ret []int
	for i, j := 0, n; i<n && j<2*n; i, j = i+1, j+1{
		ret = append(ret, nums[i], nums[j])
	}
	return ret
}
