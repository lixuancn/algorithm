package main

import "fmt"

//464. 我能赢吗

//这种最简单直观的方法是无效的，比如（4，6） A选4，B选3，返回false
//但是，如果A选1，B选任何数，A再选任何数，A赢
//func canIWin(maxChoosableInteger int, desiredTotal int) bool {
//	n := maxChoosableInteger
//	sum := 0
//	for i := 0; i < n; i++ {
//		val := n - i
//		sum += val
//		if sum > desiredTotal {
//			if i&1 == 0 {
//				return true
//			}
//			return false
//		}
//	}
//	return false
//}

//记忆化递归
//标记已使用时，用一个int来表示，int中第几位是1，就表示这个数字用过了。题目说数字不超过20，int32位足够了。
//比如int=0=0，就是所有数字都没被用
//比如int=1=1，就是数字0被用了
//比如int=2=10，就是数字1被用了
//比如int=4=100，就是数字2被用了，其他都没用
//比如int=18=10010，就是数字1和4用了。

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger >= desiredTotal {
		return true
	}
	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}
	//记录之前已经算过的结果，比如[1,2]和[2,1]，无论后面出什么，结果是一样的，所以就可以跳过了
	dp := make(map[int]bool)
	//抛开谁先手谁后手，这个函数的意义是，当前做选择的玩家是否会赢（当前玩家就是这个子问题的先手）
	var dfs func(sum, state int) bool
	dfs = func(sum, state int) bool {
		if _, ok := dp[state]; ok {
			return dp[state]
		}
		for i := 1; i <= maxChoosableInteger; i++ {
			//判断这个数字是否被用过
			if state>>i&1 == 1 {
				continue
			}
			if sum+i >= desiredTotal {
				dp[state] = true
				return true
			}
			//当前玩家选择了i之后，对手会不会输
			if !dfs(sum+i, (1<<i)|state) {
				dp[state] = true
				return true
			}
		}
		dp[state] = false
		return false
	}
	return dfs(0, 0)
}

func main() {
	fmt.Println(canIWin(3, 6))   //true
	fmt.Println(canIWin(10, 11)) //false
	fmt.Println(canIWin(10, 0))  //true
	fmt.Println(canIWin(10, 1))  //true
	fmt.Println(canIWin(4, 6))   //true 4 3 2 1
	fmt.Println(canIWin(10, 40)) //false
}
