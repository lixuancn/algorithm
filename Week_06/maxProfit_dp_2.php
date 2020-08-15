<?php
//优化DP1 子任务：第i天卖出的话，找第i天之前的最低点作为买入点。
class Solution {

    /**
     * @param Integer[] $prices
     * @return Integer
     */
    function maxProfit($prices) {
        if(!$prices || count($prices) < 2){
            return 0;
        }
        $minBuy = null;
        $maxProfit = 0;
        foreach($prices as $price){
            if(is_null($minBuy)){
                $minBuy = $price;
            }
            //今天是最小值
            if($minBuy > $price){
                $minBuy = $price;
                continue;
            }
            //今天不是最小值
            $maxProfit = max($maxProfit, $price - $minBuy);
        }
        return $maxProfit;
    }
}
$ret = (new Solution())->maxProfit(
    [7,1,5,3,6,4]
//    [7,6,4,3,1]
);
var_dump($ret);
