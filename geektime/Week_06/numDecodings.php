<?php
//DP
class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function numDecodings($s) {
        if(!strlen($s) || $s[0] == 0){
            return 0;
        }
        //pre=dp[i-2]   cur=dp[i-1]
        $pre = $cur = 1;
        for($i=1; $i<strlen($s); $i++){
            $tmp = $cur;
            if($s[$i] == 0){
                if($s[$i-1] == 1 || $s[$i-1] == 2){
                    $cur = $pre;
                }else{
                    return 0;
                }
            }else if($s[$i-1] == '1' || ($s[$i - 1] == '2' && $s[$i] >= '1' && $s[$i] <= '6'))//十几到26之间
            {
                $cur = $cur + $pre;
            }
            $pre = $tmp;
        }
        return $cur;
    }
}
$ret = (new Solution())->numDecodings(
//    "12"
    "226"
);
var_dump($ret);
