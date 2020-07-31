<?php
//快速幂 - 迭代

class Solution {

    /**
     * @param Float $x
     * @param Integer $n
     * @return Float
     */
    function myPow($x, $n) {
        if($n == 0){
            return 1;
        }
        if($n == 1){
            return $x;
        }
        return $n >= 0 ? $this->quickMul($x, $n) : 1/$this->quickMul($x, -$n);
    }

    function quickMul($x, $n){
        $ans = 1;
        $contribute = $x;
        while($n > 0){
            if($n % 2 == 1){
                $ans *= $contribute;
            }
            $contribute *= $contribute;
            $n /= 2;
        }
        return $ans;
    }
}

$ret = (new Solution())->myPow(2, 9);
var_dump($ret);
