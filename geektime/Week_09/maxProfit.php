<?php
//DP
class Solution {
    /**
     * @param Integer[] $prices
     * @return Integer
     */
    function maxProfit($prices) {
        if(!$prices) return 0;
        $m = count($prices);
        $dp = array();
        $dp[0] = 0;
        $minBuy = $prices[0];
        for($i=1; $i<$m; $i++){
            $dp[$i] = max($dp[$i-1], $prices[$i]-$minBuy);
            $minBuy = min($prices[$i], $minBuy);
        }
        return $dp[$m-1];
    }
}
$obj = new Solution();
$r = $obj->maxProfit(
//    [7,6,4,3,1]
    [7,1,5,3,6,4]
);
var_dump($r);
