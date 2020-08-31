<?php
//DP
class Solution {
    /**
     * @param Integer[][] $obstacleGrid
     * @return Integer
     */
    function uniquePathsWithObstacles($obstacleGrid) {
        if(!$obstacleGrid || !count($obstacleGrid[0])) return 0;
        $m = count($obstacleGrid);
        $n = count($obstacleGrid[0]);
        if($obstacleGrid[0][0] == 1 || $obstacleGrid[$m-1][$n-1] == 1) return 0;
        if($m == 1 && $n == 1) return $obstacleGrid[0][0] == 1 ? 0 : 1;
        $dp = array();
        $dp[0][0] = 0;
        for($i=1; $i<$m && $obstacleGrid[$i][0]!=1; $i++) $dp[$i][0] = 1;
        for($i=1; $i<$n && $obstacleGrid[0][$i]!=1; $i++) $dp[0][$i] = 1;
        for($i=1; $i<$m; $i++){
            for($j=1; $j<$n; $j++) {
                var_dump($i.'-'.$j);
                $dp[$i][$j] = $obstacleGrid[$i][$j] == 1 ? 0 : (intval($dp[$i][$j-1]) + intval($dp[$i-1][$j]));
            }
        }
        return $dp[$m-1][$n-1];
    }
}
$obj = new Solution();
$r = $obj->uniquePathsWithObstacles(
    [[0]]
);
var_dump($r);
