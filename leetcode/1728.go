package main

import "fmt"

//1728. 猫和老鼠 II

//dfs + 记忆化搜索
//从题解来看，本题最优为拓扑排序，但是我还没看
func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	n, m := len(grid), len(grid[0])
	maxStep := n*m + 2
	//找出猫，鼠的位置
	catStartX, catStartY, mouseStartX, mouseStartY := 0, 0, 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 'M' {
				mouseStartX, mouseStartY = i, j
			} else if grid[i][j] == 'C' {
				catStartX, catStartY = i, j
			}
		}
	}
	//dp
	dp := [8][8][8][8][100]int{}
	var dfs func(catX, catY, mouseX, mouseY, step int) int
	//检查边界 超出返回true
	borderCheck := func(x, y int) bool {
		return x >= n || x < 0 || y >= m || y < 0
	}
	dfs = func(catX, catY, mouseX, mouseY, step int) int {
		//fmt.Println(catX, catY, mouseX, mouseY, step)
		if res := dp[catX][catY][mouseX][mouseY][step]; res != 0 {
			return res
		}
		res := -1
		//step从0开始，0是老鼠走，1是猫走
		if step&1 == 1 {
			//可以走的位置
			tempPosition := [][]int{{catX, catY}} //原地不动
			for _, dir := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
				for i := 1; i <= catJump; i++ {
					newX, newY := catX+dir[0]*i, catY+dir[1]*i
					if borderCheck(newX, newY) {
						break
					}
					if grid[newX][newY] == '#' {
						break
					}
					if newX == mouseX && newY == mouseY {
						res = 1
						break
					}
					if grid[newX][newY] == 'F' {
						res = 1
						break
					}
					tempPosition = append(tempPosition, []int{newX, newY})
				}
				//只要有赢得方案，就可以跳出循环了
				if res == 1 {
					break
				}
			}
			//还没有找到赢的方案，就看对方会不会输
			if res == -1 {
				for _, pos := range tempPosition {
					if dfs(pos[0], pos[1], mouseX, mouseY, step+1) == -1 {
						res = 1
						break
					}
				}
			}
		} else if step <= maxStep {
			tempPosition := [][]int{{mouseX, mouseY}} //原地不动
			for _, dir := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
				for i := 1; i <= mouseJump; i++ {
					newX, newY := mouseX+dir[0]*i, mouseY+dir[1]*i
					if borderCheck(newX, newY) {
						break
					}
					if grid[newX][newY] == '#' {
						break
					}
					if newX == catX && newY == catY {
						continue
					}
					if grid[newX][newY] == 'F' {
						res = 1
						break
					}
					tempPosition = append(tempPosition, []int{newX, newY})
				}
				if res == 1 {
					break
				}
			}
			if res == -1 {
				for _, pos := range tempPosition {
					if dfs(catX, catY, pos[0], pos[1], step+1) == -1 {
						res = 1
						break
					}
				}
			}
		}
		dp[catX][catY][mouseX][mouseY][step] = res
		return res
	}
	return dfs(catStartX, catStartY, mouseStartX, mouseStartY, 0) == 1
}

func main() {
	//fmt.Println(canMouseWin([]string{"####F", "#C...", "M...."}, 1, 2))                   //true
	//fmt.Println(canMouseWin([]string{"M.C...F"}, 1, 4))                                   //true
	//fmt.Println(canMouseWin([]string{"M.C...F"}, 1, 3))                                   //false
	//fmt.Println(canMouseWin([]string{"C...#", "...#F", "....#", "M...."}, 2, 5))          //false
	//fmt.Println(canMouseWin([]string{".M...", "..#..", "#..#.", "C#.#.", "...#F"}, 3, 1)) //true
	fmt.Println(canMouseWin([]string{"CM......", "#######.", "........", ".#######", "........", "#######.", "........", "F#######"}, 1, 1)) //true
}
