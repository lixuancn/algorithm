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
        //0是现金，1是股票
        $dp[0] = array(0=>0, 1=>-$prices[0]);
        for($i=1; $i<count($prices); $i++){
            $dp[$i][0] = max($dp[$i-1][0], $dp[$i-1][1] + $prices[$i]);
            $dp[$i][1] = max($dp[$i-1][1], $dp[$i-1][0] - $prices[$i]);
        }
        return $dp[count($prices)-1][0];
    }
}
$ret = (new Solution())->maxProfit(
    [7,1,5,3,6,4]
//    [7,6,4,3,1]
//    [2,4,1]
);
var_dump($ret);
