<?php
class Solution {
    /**
     * @param String $s
     * @return Integer
     */
    function numDecodings($s) {
        if(!$s || $s[0]==0) return '0';
        $dp = array();
        $dp[-1] = 1;
        $dp[0] = 1;
        for($i=1; $i<strlen($s); $i++){
            if($s[$i] == 0 && $s[$i-1] != 1 && $s[$i-1] !=2){
                return 0;
            }
            if($s[$i] != 0){
                $dp[$i] = $dp[$i-1];
            }
            if($s[$i-1] == 1 || ($s[$i-1] == 2 && $s[$i]>=0 && $s[$i] <= 6)){
                $dp[$i] += $dp[$i-2];
            }
        }
        return $dp[strlen($s)-1];
    }
}
$obj = new Solution();
$r = $obj->numDecodings(
//    '12'
//    '226'
    '11911'
);
var_dump($r);
