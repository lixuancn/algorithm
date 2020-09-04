<?php

class Solution {
    /**
     * @param String[] $s
     * @return NULL
     */
    function reverseString(&$s) {
        if(!$s) return;
        for($i=0,$j = count($s)-1; $i<$j; $i++,$j--){
            $tmp = $s[$i];
            $s[$i] = $s[$j];
            $s[$j] = $tmp;
        }
    }
}
$obj = new Solution();
$a = ["H","a","n","n","a","h"];
$r = $obj->reverseString(
    $a
);
var_dump($a);
var_dump($r);
