<?php
//动态规划
class Solution {


    /**
     * @param Integer $m
     * @param Integer $n
     * @return Integer
     */
    function uniquePaths($m, $n) {
        $ret = array_fill(0, $n, 1);
        for($i=1; $i<$m; $i++){
            for($j=1; $j<$n; $j++){
                $ret[$j] += $ret[$j-1];
            }
        }
        return $ret[$n-1];
    }
}
$ret = (new Solution())->uniquePaths(
    3,7
);
var_dump($ret);
