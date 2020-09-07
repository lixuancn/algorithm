<?php
//滑动窗口
class Solution {
    /**
     * @param String $s
     * @param String $t
     * @return String
     */
    function minWindow($s, $t) {
        $need = array();
        $m = strlen($s);
        $n = strlen($t);
        for($i=0; $i<$n; $i++) {
            $need[$t[$i]]++;
        }
        //左开右闭区间
        $left = $right = 0;
        $window = array();
        $valid = 0;
        $start = 0;
        $len = null;
        while($right < $m){
            //c 是将移入窗口的字符
            $c = $s[$right];
            //右移
            $right++;
            if(isset($need[$c])){
                $window[$c]++;
                if($window[$c] == $need[$c]){
                    $valid++;
                }
            }
            //判断左侧窗口是否要收缩
            while($valid == count($need)){
                if(is_null($len) || $right - $left < $len){
                    $start = $left;
                    $len = $right - $left;
                }
                // d 是将移出窗口的字符
                $d = $s[$left];
                //左移
                $left++;
                if(isset($need[$d])){
                    if($need[$d] == $window[$d]){
                        $valid--;
                    }
                    $window[$d]--;
                }
            }
        }
        return is_null($len) ? "" : substr($s, $start, $len);
    }
}
$obj = new Solution();
$r = $obj->minWindow(
    "ADOBECODEBANC", "ABC"
);
var_dump($r);
