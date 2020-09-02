<?php
//DP
class Solution {
    /**
     * @param Integer[] $cost
     * @return Integer
     */
    function minCostClimbingStairs($cost) {
        if(!$cost) return 0;
        $n = count($cost);
        $dp = array();
        $dp[0] = $cost[0];
        $dp[1] = $cost[1];
        for($i=2; $i<=$n; $i++){
            $c = isset($cost[$i])?$cost[$i]:0;
            $dp[$i] = min($dp[$i-1], $dp[$i-2]) + $c;
        }
        return $dp[$n];
    }
}
$obj = new Solution();
$r = $obj->minCostClimbingStairs(
//    [10, 15, 20]
    [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
);
var_dump($r);
