<?php
//贪心
class Solution {

    /**
     * @param Integer[] $prices
     * @return Integer
     */
    function maxProfit($prices) {
        $ret = 0;
        foreach ($prices as $k => $price){
            if(isset($prices[$k+1]) && $price < $prices[$k+1]){
                $ret += $prices[$k+1] - $price;
            }
        }
        return $ret;
    }
}

$ret = (new Solution())->maxProfit(
    [7,1,5,3,6,4]
);
var_dump($ret);
