//最终题解我没看懂
func minChanges(nums []int, k int) int {
	n := len(nums)
	max := 1 << 10
	inf := math.MaxInt32 / 2
	dp := make([]int, max)
	for i := range dp {
		dp[i] = inf
	}
	//边界
	dp[0] = 0

	for i:=0; i<k; i++{
		// 第 i 个组的哈希映射
		cnt := map[int]int{}
		size := 0
		for j := i; j < n; j += k {
			cnt[nums[j]]++
			size++
		}
		t2min := min(dp...)
		g := make([]int, max)
		for j := range g {
			g[j] = t2min
		}
		for mask := range g {
			// t1 则需要枚举 x 才能求出
			for x, cntX := range cnt {
				g[mask] = min(g[mask], dp[mask^x]-cntX)
			}
		}
		// 别忘了加上 size
		for j := range g {
			g[j] += size
		}
		dp = g
	}
	return dp[0]
}

func min(arr ... int)int{
	m := arr[0]
	for _, v := range arr{
		if m > v{
			m = v
		}
	}
	return m
}
