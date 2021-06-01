package main
import (
	"fmt"
)

func main() {
	candiesCount := []int{7,4,5,3,8}
	//[favoriteTypei, favoriteDayi, dailyCapi] 。
	queries := [][]int{{0,2,2},{4,2,4},{2,13,1000000000}}
	//queries = [][]int{{4,2,4}}

	candiesCount = []int{5,2,6,4,1}
	queries = [][]int{{3,1,2},{4,10,3},{3,10,100},{4,100,30},{1,3,1}}
	//true: 在每天吃 不超过 dailyCapi 颗糖果的前提下，你可以在第 favoriteDayi 天吃到第 favoriteTypei 类糖果
	r := canEat(candiesCount, queries)
	fmt.Println(r)
}

func canEat(candiesCount []int, queries [][]int) []bool {
	answer := make([]bool, len(queries))
	sum := make(map[int]int)
	sum[-1] = 0
	for candyType, candiesCount := range candiesCount{
		sum[candyType] = sum[candyType-1] + candiesCount
	}
	for i, query := range queries{
		favoriteTypei, favoriteDayi, dailyCapi := query[0], query[1], query[2]
		if (favoriteDayi+1)*dailyCapi > sum[favoriteTypei]-candiesCount[favoriteTypei] && (favoriteDayi+1)*1 <= sum[favoriteTypei]{
			answer[i] = true
		}
	}
	return answer
}

