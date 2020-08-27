<?php
//位运算
class Solution {
    /**
     * @param Integer $n
     * @return Integer
     */
    function hammingWeight($n) {
        $count = 0;
        while ($n){
            $n = $n & ($n - 1);
            $count++;
        }
        return $count;
    }
}

$obj = new Solution();
$r = $obj->reverseBits(
    8
);
var_dump($r);
