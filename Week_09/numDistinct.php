<?php
//DP
class Solution {
    /**
     * @param String $s
     * @param String $t
     * @return Integer
     */
    function numDistinct($s, $t) {
        if(!$s && !$t) return 1;
        $m = strlen($s);
        $n = strlen($t);
        $dp = array();
        for($i=0; $i<=$n; $i++) $dp[$i][0] = 0;
        for($i=0; $i<=$m; $i++) $dp[0][$i] = 1;
        for($i=1; $i<=$n; $i++){
            for($j=1; $j<=$m; $j++) {
                if($s[$j-1] == $t[$i-1]){
                    $dp[$i][$j] = $dp[$i-1][$j-1] + $dp[$i][$j-1];
                }else{
                    $dp[$i][$j] = $dp[$i][$j-1];
                }
            }
        }
        return $dp[$n][$m];
    }
}
$obj = new Solution();
$r = $obj->numDistinct(
    "babgbag", "bag"
);
var_dump($r);
