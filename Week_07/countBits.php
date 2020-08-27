<?php
//位运算
class Solution {
    /**
     * @param Integer $num
     * @return Integer[]
     */
    function countBits($num) {
        $ret = array(0);
        for($i=1; $i<=$num; $i++){
            $k = $i & ($i - 1);
            $ret[$i] = (isset($ret[$k]) ? $ret[$k] : 0) + 1;
        }
        return $ret;
    }
}

$obj = new Solution();
$r = $obj->countBits(
    2
);
var_dump($r);
