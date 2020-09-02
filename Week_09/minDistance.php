<?php
//DP
class Solution {
    /**
     * @param String $word1
     * @param String $word2
     * @return Integer
     */
    function minDistance($word1, $word2) {
        $n1 = strlen($word1);
        $n2 = strlen($word2);
        $dp = array();
        $dp[0][0] = 0;
        //把n1变为空串的步骤
        for($i=1; $i<=$n1; $i++) $dp[$i][0] = $dp[$i-1][0] + 1;
        //把空串变成n2的步骤
        for($i=1; $i<=$n2; $i++) $dp[0][$i] = $dp[0][$i-1] + 1;
        for($i=1; $i<=$n1; $i++){
            for($j=1; $j<=$n2; $j++) {
                if($word1[$i-1] == $word2[$j-1]){
                    $dp[$i][$j] = $dp[$i-1][$j-1];
                }else{
                    $dp[$i][$j] = 1 + min($dp[$i-1][$j-1], $dp[$i-1][$j], $dp[$i][$j-1]);
                }
            }
        }
        var_dump($dp);
        return $dp[$n1][$n2];
    }
}
$obj = new Solution();
$r = $obj->minDistance(
    'horse', 'ros'
);
var_dump($r);
