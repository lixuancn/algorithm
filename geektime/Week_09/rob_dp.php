<?php
//DP
class Solution {
    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function rob1($nums) {
        if(!$nums) return 0;
        $n = count($nums);
        $dp = array();
        $dp[0][0] = 0;
        $dp[0][1] = $nums[0];
        for($i=1; $i<$n; $i++){
            $dp[$i][0] = max($dp[$i-1][0], $dp[$i-1][1]);
            $dp[$i][1] = $dp[$i-1][0] + $nums[$i];
        }
        return max($dp[$n-1][0], $dp[$n-1][1]);
    }

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function rob($nums) {
        if(!$nums) return 0;
        $n = count($nums);
        $dp = array();
        $dp[0] = $nums[0];
        $dp[1] = max($nums[0], $nums[1]);
        for($i=1; $i<$n; $i++){
            $dp[$i] = max($dp[$i-1], $dp[$i-2] + $nums[$i]);
        }
        return $dp[$n-1];
    }
}
$obj = new Solution();
$r = $obj->rob(
    [2,7,9,3,1]
);
var_dump($r);
