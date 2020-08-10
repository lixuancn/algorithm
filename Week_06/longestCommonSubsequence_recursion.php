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
        return $this->helper($m-1, $n-1, $text1, $text2, $m, $n);
    }

    private function helper($i, $j, $text1, $text2, $m, $n){
        if($i <= -1 || $j <= -1){
            return 0;
        }
        if($text1[$i] == $text2[$j]){
            return 1 + $this->helper($i-1, $j-1, $text1, $text2, $m, $n);
        }else{
            $r1 = $this->helper($i-1, $j, $text1, $text2, $m, $n);
            $r2 = $this->helper($i, $j-1, $text1, $text2, $m, $n);
            return max($r1, $r2);
        }
    }
}
$ret = (new Solution())->longestCommonSubsequence(
    "abcde", "ace"
);
var_dump($ret);

