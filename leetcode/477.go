package main
import (
	"fmt"
)

func main() {
	r := totalHammingDistance([]int{4,14,4})
	fmt.Println(r)
}

//长度为n的数组，数组中的每个数，都转成二进制后，二进制的第i位上，有c个1，n-c个0。
//所以汉明距离就是c * (n-1)
//白话：某一位上只有 1 和 0 之间有距离，1 和 1 以及 0 和 0之间没有距离
// 所以只要用数组各个数同一位上 1 的数量乘以这一位上 0 的数量结果就是这一位上的汉明距离
// 把 32 位都计算一遍然后相加即可
//由于数组中元素的范围为从 0到 10^9。 10^9 < 2^30次方，所以循环只需要i<30即可
func totalHammingDistance(nums []int) int {
	ret := 0
	n := len(nums)
	for i:=0; i<30; i++{
		c := 0
		for j := range nums{
			c += nums[j] >> i & 1
		}
		ret += c * (n-c)
	}
	return ret
}

func totalHammingDistance_violence(nums []int) int {
	ret := 0
	n := len(nums)
	for i:=0; i<n; i++{
		for j:=i+1; j<n; j++ {
			ret += getHammingDistance(nums[i], nums[j])
		}
	}
	return ret
}

//Brian Kernighan算法
func getHammingDistance(x, y int)int{
	count := 0
	for s:=x^y; s>0; s&=s-1{
		count++
	}
	return count
}

//自建轮子法
func getHammingDistance_by_self(x, y int)int{
	count := 0
	for s:=x^y; s>0; s>>=1{
		//判断最后一位是否为1，如果是1，则计数器加1
		count += s & 1
	}
	return count
}
