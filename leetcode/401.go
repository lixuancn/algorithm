package main

import (
	"fmt"
	"math/bits"
)

func main() {
	r := readBinaryWatch(1)
	fmt.Println(r)
}

//枚举二进制
//时间空间都是O(1)
//枚举所有的1024中组合(2^10)。高4位表示小时，低6位表示分钟。
func readBinaryWatch(turnedOn int) []string {
	ans := make([]string, 0)
	for i:=0; i<1024; i++{
		// 用位运算取出高 4 位和低 6 位
		h, m := i>>6, i&63
		if h < 12 && m < 60 && bits.OnesCount(uint(i)) == turnedOn{
			ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
		}
	}
	return ans
}

//枚举时分
//时间空间都是O(1)。枚举次数和传入的值无关。
func readBinaryWatch_loop(turnedOn int) []string {
	ans := make([]string, 0)
	for h:=uint8(0); h<12; h++{
		for m:=uint8(0); m<60; m++ {
			if bits.OnesCount8(h) + bits.OnesCount8(m) == turnedOn{
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return ans
}

//递归
var valueList = []int{8,4,2,1,32,16,8,4,2,1}
var ret []string
func readBinaryWatch_recursion(turnedOn int) []string {
	ret = make([]string, 0)
	recursion(turnedOn, 0, 0, 0, 0)
	return ret
}

func recursion(turnedOn, index, n, hour, minute int){
	if hour > 11 || minute > 59{
		return
	}
	if turnedOn == n{
		t := fmt.Sprintf("%d:%02d", hour, minute)
		ret = append(ret, t)
		return
	}
	if len(valueList) == index{
		return
	}
	var hour1, minute1, hour2, minute2 = hour, minute, hour, minute
	if index < 4{
		hour1 = hour + valueList[index]
		hour2 = hour
	}else{
		minute1 = minute + valueList[index]
		minute2 = minute
	}
	recursion(turnedOn, index+1, n+1, hour1, minute1)
	recursion(turnedOn, index+1, n, hour2, minute2)
}
