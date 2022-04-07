<?php

class Solution {
    /**
     * @param String $s
     * @param Integer $k
     * @return String
     */
    function reverseStr($s, $k) {
        if(!$s) return $s;
        $n = strlen($s);
        $start = 0;
        while (1){
            $end = ($n>$start+$k ? $start+$k : $n) - 1;
            for($i=$start,$j=$end; $i<$j; $i++,$j--){
                $tmp = $s[$i];
                $s[$i] = $s[$j];
                $s[$j] = $tmp;
            }
            $start = $start + 2*$k;
            if($start >= $n-1){
                break;
            }
        }
        return $s;
    }
}
$obj = new Solution();
$r = $obj->reverseStr(
    'abcdefg', 2
);
var_dump($r);
