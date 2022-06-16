package main

import (
	"fmt"
	"sort"
)

//675. 为高尔夫比赛砍树

func cutOffTree(forest [][]int) int {
	//return cutOffTree_backtracking(forest) //DFS 会超时
	return cutOffTree_bfs(forest) //bfs
}

//回溯 dfs 自己想的
//思路是求每一步的最小距离，即树0到树1的最短距离 + 树1到树2的最短距离
func cutOffTree_backtracking(forest [][]int) int {
	posList := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	m, n := len(forest), len(forest[0])
	treeList := make([]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if forest[i][j] > 1 {
				treeList = append(treeList, forest[i][j])
			}
		}
	}
	if len(treeList) <= 0 {
		return -1
	}
	sort.Ints(treeList)
	result, ret := 0, 1<<31
	startX, startY := 0, 0
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}

	var backtracking func(i, j, target, step int)
	backtracking = func(i, j, target, step int) {
		//已经超过之前的结果了，剪掉
		if step > ret {
			return
		}
		if forest[i][j] == target {
			if ret > step {
				startX, startY = i, j
				ret = step
			}
			return
		}
		for _, pos := range posList {
			x, y := i+pos[0], j+pos[1]
			if x < 0 || x >= m || y < 0 || y >= n || forest[x][y] == 0 {
				continue
			}
			if used[x][y] {
				continue
			}
			used[x][y] = true
			backtracking(x, y, target, step+1)
			used[x][y] = false
		}
	}
	for _, target := range treeList {
		ret = 1 << 31
		used[startX][startY] = true
		oldX, oldY := startX, startY
		backtracking(startX, startY, target, 0)
		used[oldX][oldY] = false
		//没找到
		if ret == 1<<31 {
			return -1
		}
		result += ret
	}
	return result
}

//bfs
var dir4 = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func cutOffTree_bfs(forest [][]int) (ans int) {
	type pair struct{ dis, x, y int }
	trees := []pair{}
	for i, row := range forest {
		for j, h := range row {
			if h > 1 {
				trees = append(trees, pair{h, i, j})
			}
		}
	}
	sort.Slice(trees, func(i, j int) bool { return trees[i].dis < trees[j].dis })

	bfs := func(sx, sy, tx, ty int) int {
		m, n := len(forest), len(forest[0])
		vis := make([][]bool, m)
		for i := range vis {
			vis[i] = make([]bool, n)
		}
		vis[sx][sy] = true
		q := []pair{{0, sx, sy}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			if p.x == tx && p.y == ty {
				return p.dis
			}
			for _, d := range dir4 {
				if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && !vis[x][y] && forest[x][y] > 0 {
					vis[x][y] = true
					q = append(q, pair{p.dis + 1, x, y})
				}
			}
		}
		return -1
	}

	preX, preY := 0, 0
	for _, t := range trees {
		d := bfs(preX, preY, t.x, t.y)
		if d < 0 {
			return -1
		}
		ans += d
		preX, preY = t.x, t.y
	}
	return
}

func main() {
	fmt.Println(cutOffTree([][]int{{1, 2, 3}, {0, 0, 4}, {7, 6, 5}})) //6
	fmt.Println(cutOffTree([][]int{{1, 2, 3}, {0, 0, 0}, {7, 6, 5}})) //-1
	fmt.Println(cutOffTree([][]int{{2, 3, 4}, {0, 0, 5}, {8, 7, 6}})) //6
}
