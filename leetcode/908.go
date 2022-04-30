package main

import (
	"fmt"
)

//908. 最小差值 I

func smallestRangeI(nums []int, k int) int {
	//找出最大最小值，然后求这两个数的中位数（都往中间靠），这样分数最小。
	maxDigint, minDigint := -1<<31, 1<<31
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxDigint {
			maxDigint = nums[i]
		}
		if nums[i] < minDigint {
			minDigint = nums[i]
		}
	}
	//下面这段是我自己想的，但是题解更精简，是数学法。就一行：ans := maxNum - minNum - 2*k   ans>0?ans:0
	//我也想明白了，如果max-min <= 2k，那么max-x，min+x之后，max=min，分数为0。数组里的所有元素都可以变成同一个相同元素
	//如果max-min>2k，那么最小分数就是(max-k) - (min+k) = max-min-2k
	median := (maxDigint + minDigint) >> 1
	if minDigint+k >= median {
		minDigint = median
	} else {
		minDigint += k
	}
	if maxDigint-k <= median {
		maxDigint = median
	} else {
		maxDigint -= k
	}
	return maxDigint - minDigint
}

func main() {
	fmt.Println(smallestRangeI([]int{1}, 0))       //0
	fmt.Println(smallestRangeI([]int{0, 10}, 2))   //6
	fmt.Println(smallestRangeI([]int{1, 3, 6}, 3)) //0
}
