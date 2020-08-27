<?php
//位运算
class Solution {
    /**
     * @param Integer $n
     * @return Integer
     */
    function reverseBits($n) {
        $power = 31;
        $ret = 0;
        while ($n){
            $ret += ($n & 1) << $power;
            $n = $n >> 1;
            $power--;
        }
        return $ret;
    }
}

//964176192
//var_dump(bindec('00111001011110000010100101000000'));
$obj = new Solution();
$r = $obj->reverseBits(
    //43261596
    bindec('00000010100101000001111010011100')
);
var_dump($r);
