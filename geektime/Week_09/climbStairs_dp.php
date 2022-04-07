<?php
//DP
class Solution {
    /**
     * @param Integer $n
     * @return Integer
     */
    function climbStairs($n) {
        if($n == 0) return 0;
        if($n == 1) return 1;
        if($n == 2) return 2;
        $dp = array();
        $dp[0] = 0;
        $dp[1] = 1;
        $dp[2] = 2;
        for($i=3; $i<=$n; $i++){
            $dp[$i] = $dp[$i-1] + $dp[$i-2];
        }
        return $dp[$n];
    }
}
$obj = new Solution();
$r = $obj->climbStairs(
    3
);
var_dump($r);
