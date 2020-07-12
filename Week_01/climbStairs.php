<?php

/**
 * 动态规划，最后一阶，可能跨一步也可能跨两步，那么第N阶=f(n) = f(n-1) + f(n-2)
 * 数学归纳法也可以证明：0+1=1  1+2=3 2+3=5
 * Class Solution
 */
class Solution {

    /**
     * @param Integer $n
     * @return Integer
     */
    function climbStairs($n) {
        $p = $q = 0;
        $r = 1;
        for($i=1; $i<=$n; $i++){
            $p = $q;
            $q = $r;
            $r = $p + $q;
        }
        return $r;
    }
}
$obj = new Solution();
$nums = $obj->climbStairs(5);
print_r($nums);
