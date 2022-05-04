package main

import "fmt"

//1823. 找出游戏的获胜者

func findTheWinner(n int, k int) int {
	//return findTheWinner_queue(n, k)       //用队列模拟
	return findTheWinner_yuesefuhuan(n, k) //约瑟夫环
}

//模拟，出队在入队，如果是淘汰的则不入队
func findTheWinner_queue(n int, k int) int {
	queue := make([]int, n)
	for i := range queue {
		queue[i] = i + 1
	}
	for len(queue) > 1 {
		for i := 1; i < k; i++ {
			//出队
			val := queue[0]
			queue = queue[1:]
			//入队
			queue = append(queue, val)
		}
		//需要淘汰
		queue = queue[1:]
	}
	return queue[0]
}

//约瑟夫环的经典问题。
//由于本次起点，和下次的起点，在原位置中相距k，那么本次问题的答案就转为findTheWinner(n-1, k)+k的问题
func findTheWinner_yuesefuhuan(n int, k int) int {
	if n == 1 {
		return 1
	}
	ans := (findTheWinner(n-1, k) + k) % n
	if ans == 0 {
		return n
	}
	return ans
}

func main() {
	fmt.Println(findTheWinner(5, 2)) //3
	fmt.Println(findTheWinner(6, 5)) //1
}
