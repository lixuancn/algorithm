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
        $dp = array_fill(0, count($prices), 0);
        $minBuy = $prices[0];
        for($i=1; $i<count($prices); $i++){
            $minBuy = min($prices[$i], $minBuy);
            $dp[$i] = max($dp[$i-1], $prices[$i] - $minBuy);
        }
        return isset($dp[count($prices)-1]) ? $dp[count($prices)-1] : 0;
    }
}
$ret = (new Solution())->maxProfit(
//    [7,1,5,3,6,4]
//    [7,6,4,3,1]
    [2,4,1]
);
var_dump($ret);
