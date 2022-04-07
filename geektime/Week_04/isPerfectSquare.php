<?php

//牛顿迭代法。
class Solution {

    /**
     * @param Integer $num
     * @return Boolean
     */
    function isPerfectSquare($num) {
        if($num == 0 || $num == 1){
            return $num;
        }
        $r = $num;
        while(1){
            if($r * $r <= $num){
                break;
            }
            //这是数学证明的公式
            $tmp = ($r + $num/$r) / 2;
            if ($r - $tmp < 1e-7) {
                break;
            }
            $r = $tmp;
        }
        $ret = intval($r);
        if($ret * $ret == $num){
            return true;
        }
        return false;
    }
}
