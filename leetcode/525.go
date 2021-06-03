package main
import (
	"fmt"
)

func main() {
	r := findMaxLength([]int{0,0,1,0,0,0,1,1})
	fmt.Println(r)
}

//前缀和
//把1当做1，把0当做-1，这样的话0和1个数相当于的子数组和为0
func findMaxLength(nums []int) int {
	sum, max := 0, 0
	var preSum = map[int]int{0: -1}
	for i, num := range nums{
		if num == 0{
			sum--
		}else{
			sum++
		}
		if j, ok := preSum[sum]; ok{
			if max < i-j{
				max = i - j
			}
		}else{
			preSum[sum] = i
		}
	}
	return max
}


//递归+剪枝
var cache = map[string]int{}
func findMaxLength_recursion(nums []int) int {
	cache = map[string]int{}
	zeroCount, oneCount := 0, 0
	for _, num := range nums{
		if num == 0{
			zeroCount++
		}else{
			oneCount++
		}
	}
	return recursion(nums, len(nums), zeroCount, oneCount, 0, len(nums)-1)
}

func recursion(nums []int, n, zeroCount, oneCount, offsetMin, offsetMax int)int{
	if n < 2{
		return 0
	}
	if zeroCount == oneCount{
		return n
	}
	cacheKey := fmt.Sprintf("%d-%d", offsetMin, offsetMax)
	if _, ok := cache[cacheKey]; ok{
		return cache[cacheKey]
	}
	newZeroCount, newOneCount := zeroCount, oneCount
	if nums[n-1] == 0{
		newZeroCount = zeroCount - 1
	}else{
		newOneCount = oneCount - 1
	}
	r1 := recursion(nums[0:n-1], n-1, newZeroCount, newOneCount, offsetMin, offsetMax-1)

	newZeroCount, newOneCount = zeroCount, oneCount
	if nums[0] == 0{
		newZeroCount = zeroCount - 1
	}else{
		newOneCount = oneCount - 1
	}
	r2 := recursion(nums[1:], n-1, newZeroCount, newOneCount, offsetMin+1, offsetMax)

	max := r1
	if r2 > r1{
		max = r2
	}
	cache[cacheKey] = max
	return max
}
