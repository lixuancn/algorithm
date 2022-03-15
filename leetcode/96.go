package main

func numTrees(n int) int {
	dp := make([]int, n+1)
	//空树也是一个BST也是
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			//总共有i个节点，以j为节点时，左子树有j-1个节点，右子树有i-j个节点
			//遍历j后需要加总
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
