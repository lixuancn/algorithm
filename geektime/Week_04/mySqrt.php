<?php
//二分.
//原因：y = x^2, 1单调递增的，2有上下界（0和x）
class Solution {

    /**
     * @param Integer $x
     * @return Integer
     */
    function mySqrt($x) {
        if($x == 0 || $x == 1){
            return $x;
        }
        $left = 0;
        $right = $x;
        $ans = 0;
        while($left <= $right){
            $mid = $left + ($right - $left) / 2;
            $mid = intval($mid);
            if($mid * $mid <= $x){
                $ans = $mid;
                $left = $mid + 1;
            }else{
                $right = $mid - 1;
            }
        }
        return intval($ans);
    }
}

$ret = (new Solution())->mySqrt(
    5
);
var_dump($ret);
