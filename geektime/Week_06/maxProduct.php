<?php
//dp
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function maxProduct($nums) {
        $ans = $max = $min = $nums[0];
        for($i=1; $i<count($nums); $i++){
            $mx = $max;
            $mn = $min;
            $max = max($mx * $nums[$i], $nums[$i], $mn * $nums[$i]);
            $min = min($mn * $nums[$i], $nums[$i], $mx * $nums[$i]);
            $ans = max($max, $ans);
        }
        return $ans;
    }
}
$ret = (new Solution())->maxProduct(
//    [2,3,-2,4]
//    [-2,0,-1]
    [0,2]
);
var_dump($ret);
