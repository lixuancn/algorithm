package main

import (
	"fmt"
)

//1109. 航班预订统计

//这里有 n 个航班，它们分别从 1 到 n 进行编号。
//有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi]
//意味着在从 firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。
//请你返回一个长度为 n 的数组 answer，其中 answer[i] 是航班 i 上预订的座位总数。

func main() {
	var bookings = [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
	//输出：[10,55,45,25,25]
	//航班编号        1   2   3   4   5
	//预订记录 1 ：   10  10
	//预订记录 2 ：       20  20
	//预订记录 3 ：       25  25  25  25
	//总座位数：      10  55  45  25  25
	//因此，answer = [10,55,45,25,25]

	r := corpFlightBookings_practice(bookings, 5)
	fmt.Println(r)
}

//差分法
func corpFlightBookings(bookings [][]int, n int) []int {
	//0~n+1
	details := make([]int, n+2)
	for _, booking := range bookings {
		details[booking[0]] += booking[2]
		details[booking[1]+1] -= booking[2]
	}
	var answer []int
	curr := details[0]
	for i := 1; i <= n; i++ {
		curr += details[i]
		answer = append(answer, curr)
	}
	return answer
}

//[1,3]航班被预定了10张票，那么下标0就是10，下标3就是-10，累加法
//https://mp.weixin.qq.com/s/Nga07U-FgH5fCE7TlqcByw
func corpFlightBookings_practice(bookings [][]int, n int) []int {
	result := make([]int, n+1)
	for _, booking := range bookings {
		result[booking[0]-1] += booking[2]
		result[booking[1]-1+1] += -booking[2]
	}
	for i := 1; i < len(result); i++ {
		result[i] = result[i] + result[i-1]
	}
	return result[:len(result)-1]
}
