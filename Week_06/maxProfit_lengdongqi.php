<?php
//DP
class Solution {

    /**
     * @param Integer[] $prices
     * @return Integer
     */
    function maxProfit($prices) {
        if(!$prices || count($prices) < 2){
            return 0;
        }
        $dp = array();
        //0是股票，1是现金且处于冷冻期，2是现金且处于非冷冻期
        $dp[0][0] = -$prices[0];
        $dp[0][1] = 0;
        $dp[0][2] = 0;
        for($i=1; $i<count($prices); $i++){
            $dp[$i][0] = max($dp[$i-1][0], $dp[$i-1][2] - $prices[$i]);
            $dp[$i][1] = $dp[$i-1][0] + $prices[$i];
            $dp[$i][2] = max($dp[$i-1][1], $dp[$i-1][2]);
        }
        var_dump($dp);
        return max($dp[count($prices)-1][1], $dp[count($prices)-1][2]);
    }
}
$ret = (new Solution())->maxProfit(
//    [1,2,3,0,2]
//    [1,2,4]
    [2, 1]
);
var_dump($ret);
