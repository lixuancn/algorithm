package main
import (
	"fmt"
	"math"
)

func main() {
	s := "aaabbb"

	r := strangePrinter(s)
	fmt.Println(r)
}

//动态规划
//dp[i][j] 表示从i到j最少操作数。
//如果s[i]==s[j] 则j可以通过i打印，那么dp[i][j] = dp[i][j-1]
//如果s[i]!=s[j]，则dp[i][j] = dp[i][k] + dp[k+1][j]。k是从i到j-1，循环k，取最小的dp[i][j]。即dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j])
//为了循环时每次都能拿到值（[j-1]要有值）（无后效性），则需要i从n到0，j需要从i到j
func strangePrinter(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp{
		dp[i] = make([]int, n)
	}
	for i:=n-1; i>=0; i--{
		dp[i][i] = 1
		for j:=i+1; j<n; j++{
			if s[i] == s[j]{
				dp[i][j] = dp[i][j-1]
			}else{
				dp[i][j] = math.MaxInt32
				for k:=i; k<j; k++{
					dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j])
				}
			}
		}
	}
	return dp[0][n-1]
}

func min(i, j int)int{
	if i < j{
		return i
	}
	return j
}
