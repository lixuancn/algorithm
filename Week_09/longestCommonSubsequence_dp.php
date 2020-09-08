<?php
//DP
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
        $dp = array();
        for($i=0; $i<$m; $i++){
            for($j=0; $j<$n; $j++) {
                if($text1[$i] == $text2[$j]){
                    $dp[$i][$j] = 1 + intval($dp[$i-1][$j-1]);
                }else{
                    $dp[$i][$j] = max($dp[$i][$j-1], $dp[$i-1][$j]);
                }
            }
        }
        return isset($dp[$m-1][$n-1]) ? $dp[$m-1][$n-1] : 0;
    }
}
$obj = new Solution();
$r = $obj->longestCommonSubsequence(
    "abcde", "ace"
);
var_dump($r);
