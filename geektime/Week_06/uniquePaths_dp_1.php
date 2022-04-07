<?php
//dp 
class Solution {


    /**
     * @param Integer $m
     * @param Integer $n
     * @return Integer
     */
    function uniquePaths($m, $n) {
        $ret = array_fill(0, $m, array_fill(0, $n, 1));
        for($i=1; $i<$m; $i++){
            for($j=1; $j<$n; $j++){
                $ret[$i][$j] = $ret[$i-1][$j] + $ret[$i][$j-1];
            }
        }
        return $ret[$m-1][$n-1];
    }
}
$ret = (new Solution())->uniquePaths(
    3,7
);
var_dump($ret);

