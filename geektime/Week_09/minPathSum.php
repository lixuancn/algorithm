<?php
//DP
class Solution {
    /**
     * @param Integer[][] $grid
     * @return Integer
     */
    function minPathSum($grid) {
        if(!$grid) return 0;
        $m = count($grid);
        $n = count($grid[0]);
        $dp = array();
        $dp[0][0] = $grid[0][0];
        for($i=1; $i<$m; $i++){
            $dp[$i][0] = $dp[$i-1][0] + $grid[$i][0];
        }
        for($i=1; $i<$n; $i++){
            $dp[0][$i] = $dp[0][$i-1] + $grid[0][$i];
        }
        for($i=1; $i<$m; $i++){
            for($j=1; $j<$n; $j++) {
                $dp[$i][$j] = $grid[$i][$j] + min($dp[$i-1][$j], $dp[$i][$j-1]);
            }
        }
        return $dp[$m-1][$n-1];
    }
}
$obj = new Solution();
$r = $obj->minPathSum(
    [
        [1,3,1],
        [1,5,1],
        [4,2,1]
    ]
);
var_dump($r);
