package main
import (
	"fmt"
)

func main() {
	p1 := []int{1,3}

	r := subsetXORSum(p1)
	fmt.Println(r)
}

func subsetXORSum(nums []int) int {
	return dfs(nums)
}

//递归找子集
func dfs(nums []int)int{
	var recursion func (res, val, index, n int, nums []int)int
	recursion = func (res, val, index, n int, nums []int)int{
		if index == n{
		res += val
		return res
	}
		//子集有这个数
		r1 := recursion(res, val^nums[index], index+1, n, nums)
		//子集没这个数
		r2 := recursion(res, val, index+1, n, nums)
		return r1 + r2
	}

	return recursion(0, 0, 0, len(nums), nums)
}
