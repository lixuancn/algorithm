<?php
//位运算 - 2的幂的数对应二进制一定有且只有一个1
class Solution {

    /**
     * @param Integer $n
     * @return Boolean
     */
    function isPowerOfTwo($n) {
        return $n > 0 && (($n & ($n-1)) == 0);
    }
}

$obj = new Solution();
$r = $obj->isPowerOfTwo(
    2
);
var_dump($r);
