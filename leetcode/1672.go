package main

import (
	"fmt"
)

//给你一个 m x n 的整数网格 accounts ，其中 accounts[i][j] 是第 i 位客户在第 j 家银行托管的资产数量。返回最富有客户所拥有的 资产总量 。
//客户的 资产总量 就是他们在各家银行托管的资产数量之和。最富有客户就是 资产总量 最大的客户。

func main() {
	var accounts = [][]int{{1, 2, 3}, {3, 2, 1}}
	r := maximumWealth(accounts)
	fmt.Println(r)
}

func maximumWealth(accounts [][]int) int {
	var max int
	for _, bankList := range accounts {
		var userMoney = 0
		for _, money := range bankList {
			userMoney += money
		}
		if max < userMoney {
			max = userMoney
		}
	}
	return max
}
