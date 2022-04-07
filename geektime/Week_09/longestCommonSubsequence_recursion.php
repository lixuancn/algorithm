<?php
//递归
class Solution {
    /**
     * @param String $text1
     * @param String $text2
     * @return Integer
     */
    function longestCommonSubsequence($text1, $text2) {
        $m = strlen($text1);
        $n = strlen($text2);
        if(!$m || !$n){
            return 0;
        }
        $ans = $this->helper($text1, $text2, $m, $n, 0, 0);
        return $ans;
    }

    function helper($text1, $text2, $m, $n, $i, $j){
        if($i>=$m || $j>=$n) return 0;
        if($text1[$i] == $text2[$j]){
            return 1 + $this->helper($text1, $text2, $m, $n, $i+1, $j+1);
        }else{
            return max(
                $this->helper($text1, $text2, $m, $n, $i+1, $j),
                $this->helper($text1, $text2, $m, $n, $i, $j+1)
            );
        }
    }
}
$obj = new Solution();
$r = $obj->longestCommonSubsequence(
    "abcde", "ace"
);
var_dump($r);
