<?php
//DP
class Solution {
    /**
     * @param Integer $m
     * @param Integer $n
     * @return Integer
     */
    function uniquePaths($m, $n) {
        if($m == 0) return 0;
        if($m == 1 && $n == 1) return 1;
        $dp = array();
        $dp[0][0] = 0;
        for($i=1; $i<$m; $i++) $dp[$i][0] = 1;
        for($i=1; $i<$n; $i++) $dp[0][$i] = 1;
        for($i=1; $i<$m; $i++){
            for($j=1; $j<$n; $j++) {
                $dp[$i][$j] = intval($dp[$i][$j-1]) + intval($dp[$i-1][$j]);
            }
        }
        return $dp[$m-1][$n-1];
    }
}
$obj = new Solution();
$r = $obj->uniquePaths(
    3, 2
);
var_dump($r);
