<?php
//DP
class Solution {

    /**
     * @param Integer[] $prices
     * @param Integer $fee
     * @return Integer
     */
    function maxProfit($prices, $fee) {
        if(!$prices || count($prices) < 2){
            return 0;
        }
        $dp = array();
        //0是现金，1是股票
        $dp[0] = array(0=>0, 1=>-$prices[0]);
        for($i=1; $i<count($prices); $i++){
            $dp[$i][0] = max($dp[$i-1][0], $dp[$i-1][1] + $prices[$i] - $fee);
            $dp[$i][1] = max($dp[$i-1][1], $dp[$i-1][0] - $prices[$i]);
        }
        return $dp[count($prices)-1][0];
    }
}
$ret = (new Solution())->maxProfit(
    [1, 3, 2, 8, 4, 9], 2
);
var_dump($ret);
