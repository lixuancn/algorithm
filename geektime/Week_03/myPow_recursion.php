<?php
//快速幂 - 递归

class Solution {

    /**
     * @param Float $x
     * @param Integer $n
     * @return Float
     */
    function myPow($x, $n) {
        return $n >= 0 ? $this->quickMul($x, $n) : 1/$this->quickMul($x, -$n);
    }

    function quickMul($x, $n){
        if($n == 0){
            return 1;
        }
        if($n == 1){
            return $x;
        }
        $y = $this->quickMul($x, intval($n / 2));
        //奇数偶数
        return $n % 2 != 0 ? $y * $y * $x : $y * $y;
    }
}

$ret = (new Solution())->myPow(2, 9);
var_dump($ret);
