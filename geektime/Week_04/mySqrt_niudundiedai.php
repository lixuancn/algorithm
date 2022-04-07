<?php
//牛顿迭代法。
class Solution {

    /**
     * @param Integer $x
     * @return Integer
     */
    function mySqrt($x) {
        if($x == 0 || $x == 1){
            return $x;
        }
        $r = $x;
        while(1){
            if($r * $r <= $x){
                break;
            }
            //这是数学证明的公式
            $tmp = ($r + $x/$r) / 2;
            if ($r - $tmp < 1e-7) {
                break;
            }
            $r = $tmp;
        }
        return intval($r);
    }
}

$ret = (new Solution())->mySqrt(
    5
);
var_dump($ret);
