package main

import "fmt"

//699. 掉落的方块

func fallingSquares(positions [][]int) []int {
	return fallingSquares_force(positions)       //暴力法
	return fallingSquares_segmentTree(positions) //线段树，可以优化到nlogn 还没看懂 也没跑过
}

func fallingSquares_force(positions [][]int) []int {
	n := len(positions)
	heights := make([]int, n) //每个方块落下后的高度
	for i := 0; i < n; i++ {
		left1, right1 := positions[i][0], positions[i][0]+positions[i][1]
		//落照落点，会落在哪个方块上
		maxHeight, landPos := 0, -1
		for j := 0; j < i; j++ {
			left2, right2 := positions[j][0], positions[j][0]+positions[j][1]
			if right1 <= left2 || left1 >= right2 {
				continue
			}
			if heights[j] > maxHeight {
				maxHeight = heights[j]
				landPos = j
			}
		}
		//本方块降落高度就是落点的高度加上自己
		heights[i] = positions[i][1]
		if landPos != -1 {
			heights[i] += heights[landPos]
		}
	}
	for i := 1; i < n; i++ {
		if heights[i] < heights[i-1] {
			heights[i] = heights[i-1]
		}
	}
	return heights
}

type SegmentNode struct {
	ls  *SegmentNode // 代表当前区间的左子节点
	rs  *SegmentNode // 代表当前区间的右子节点
	val int          // val 代表当前区间的最大高度
	add int          // add 为懒标记
}

type SegmentTree struct {
	root *SegmentNode
}

func NewSegmentTree() *SegmentTree {
	return &SegmentTree{root: &SegmentNode{}}
}

func (st *SegmentTree) update(node *SegmentNode, lc, rc, l, r, v int) {
	if l <= rc && rc <= r {
		node.val, node.add = v, v
		return
	}
	st.pushdown(node)
	mid := (lc + rc) >> 1
	if l <= mid {
		st.update(node.ls, lc, mid, l, r, v)
	}
	if r > mid {
		st.update(node.ls, mid+1, rc, l, r, v)
	}
	st.pushup(node)
}

func (st *SegmentTree) pushdown(node *SegmentNode) {
	if node.ls == nil {
		node.ls = &SegmentNode{}
	}
	if node.rs == nil {
		node.rs = &SegmentNode{}
	}
	if node.add == 0 {
		return
	}
	node.ls.add = node.add
	node.rs.add = node.add
	node.ls.val = node.add
	node.rs.val = node.val
	node.add = 0
}

func (st *SegmentTree) pushup(node *SegmentNode) {
	node.val = node.ls.val
	if node.rs.val > node.val {
		node.val = node.rs.val
	}
}

func (st *SegmentTree) query(node *SegmentNode, lc, rc, l, r int) int {
	if l <= lc && rc <= r {
		return node.val
	}
	st.pushdown(node)
	mid := (lc + rc) >> 1
	ans := 0
	if l <= mid {
		ans = st.query(node.ls, lc, mid, l, r)
	}
	if r > mid {
		ret := st.query(node.rs, mid+1, rc, l, r)
		if ans < ret {
			ans = ret
		}
	}
	return ans
}

func fallingSquares_segmentTree(positions [][]int) []int {
	n := len(positions)
	st := NewSegmentTree()
	result := make([]int, n)
	for i := 0; i < n; i++ {
		left, h := positions[i][0], positions[i][1]
		cur := st.query(st.root, 0, 1<<31, left, left+h-1)
		st.update(st.root, 0, 1<<31, left, left+h-1, cur+h)
		result[i] = st.root.val
	}
	return result
}

func main() {
	fmt.Println(fallingSquares([][]int{{1, 2}, {2, 3}, {6, 1}}))
	fmt.Println(fallingSquares([][]int{{100, 100}, {200, 100}}))
	fmt.Println(fallingSquares([][]int{{2, 1}, {2, 9}, {1, 8}}))
	fmt.Println(fallingSquares([][]int{{9, 7}, {1, 9}, {3, 1}}))
}
