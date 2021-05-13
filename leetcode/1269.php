<?php
//有一个长度为 arrLen 的数组，开始有一个指针在索引 0 处。
//每一步操作中，你可以将指针向左或向右移动 1 步，或者停在原地（指针不能被移动到数组范围外）。
//给你两个整数 steps 和 arrLen ，请你计算并返回：在恰好执行 steps 次操作以后，指针仍然指向索引 0 处的方案数。
//由于答案可能会很大，请返回方案数 模 10^9 + 7 后的结果。

//算数的，常见方案是动态规划。
//递归转动态规划的，递归函数有几个参数，DP就是几维度
//DP[i][j] 表示 第i步在j位置时的方案数。 0<=i<=steps 0<=j<=arrLen-1 跑太远回不来，所以0<=j<=min(arrLen-1, steps/2+1)
//边界：DP[0][0] = 1，DP[0][j] = 0
//DP[i][j] = DP[i-1][j-1] + DP[i-1][j] + DP[i-1][j+1]

//DP通常，每一行只和上一行有关，因此通常都可以优化掉空间复杂度
class Solution {
    /**
     * 二维dp改为一维，空间复杂度为O min(arrLen-1，steps/2+1)
     *
     * @param Integer $steps
     * @param Integer $arrLen
     * @return Integer
     */
    function numWays($steps, $arrLen) {
        $jMaxIndex = min($arrLen-1, $steps/2+1);
        $dp = array();
        //上一轮的，即i-1
        $dp[0] = 1;
        for($i=1; $i<=$steps; $i++){
            //本轮的
            $nextDp = array();
            for($j=0; $j<=$jMaxIndex; $j++){
                if($j-1 < 0){
                    $nextDp[$j] = $dp[$j] + $dp[$j+1];
                }else if($j+1 > $jMaxIndex){
                    $nextDp[$j] = $dp[$j-1] + $dp[$j];
                }else{
                    $nextDp[$j] = $dp[$j-1] + $dp[$j] + $dp[$j+1];
                }
                $nextDp[$j] = $nextDp[$j] % 1000000007;
            }
            $dp = $nextDp;
        }
        return $dp[0];
    }

    /**
     * @param Integer $steps
     * @param Integer $arrLen
     * @return Integer
     */
    function numWays_dp($steps, $arrLen) {
        $jMaxIndex = min($arrLen-1, $steps/2+1);
        $dp = array();
        $dp[0][0] = 1;
        for($i=1; $i<=$steps; $i++){
            for($j=0; $j<=$jMaxIndex; $j++){
                if($j-1 < 0){
                    $dp[$i][$j] = $dp[$i-1][$j] + $dp[$i-1][$j+1];
                }else if($j+1 > $jMaxIndex){
                    $dp[$i][$j] = $dp[$i-1][$j-1] + $dp[$i-1][$j];
                }else{
                    $dp[$i][$j] = $dp[$i-1][$j-1] + $dp[$i-1][$j] + $dp[$i-1][$j+1];
                }
                $dp[$i][$j] = $dp[$i][$j] % 1000000007;
            }
        }
        return $dp[$steps][0];
    }
}

$result = (new Solution())->numWays(47, 38);
var_dump($result);
//输出：4
//解释：3 步后，总共有 4 种不同的方法可以停在索引 0 处。
//向右，向左，不动
//不动，向右，向左
//向右，不动，向左
//不动，不动，不动

//pages/appoint/appoint?cb_fromsource=0330-xmy-jx
