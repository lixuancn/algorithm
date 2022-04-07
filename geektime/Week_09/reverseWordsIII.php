<?php
//反转字符串中的单词 III
class Solution {
    /**
     * @param String $s
     * @return String
     */
    function reverseWords($s) {
        if(!$s) return $s;
        $n = strlen($s);
        $start = $end = 0;
        for($i=0; $i<$n; $i++){
            if($s[$i] != ' ' && $i != $n-1){
                continue;
            }
            $end = $i == $n-1 ? $i : $i-1;
            for(; $start<$end; $start++, $end--){
                $tmp = $s[$start];
                $s[$start] = $s[$end];
                $s[$end] = $tmp;
            }
            $start = $i+1;
        }
        return $s;
    }
}
$obj = new Solution();
$r = $obj->reverseWords(
    "Let's take LeetCode contest"
);
var_dump($r);
