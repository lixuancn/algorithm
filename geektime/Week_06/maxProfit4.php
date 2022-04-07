<?php
//dp
//dp[i][j][0] = max(dp[i - 1][j][0], dp[i - 1][j][1] + prices[i])
//dp[i][j][1] = max(dp[i - 1][j][1], dp[i - 1][j - 1][0] - prices[i])
class Solution {
    /**
     * @param Integer $k
     * @param Integer[] $prices
     * @return Integer
     */
    function maxProfit($k, $prices) {
        if(!$prices || count($prices) < 1){
            return 0;
        }
        $dp = array_fill(0, count($prices), array_fill(0, $k+1, [0, 0]));
        for($i=0; $i<count($prices); $i++){
            for($j=1; $j<$k+1; $j++){
                if($i==0){
                    $dp[0][$j][1] = -$prices[$i];
                    continue;
                }
                $dp[$i][$j][0] = max($dp[$i-1][$j][0], $dp[$i-1][$j][1]+$prices[$i]);
                $dp[$i][$j][1] = max($dp[$i-1][$j][1], $dp[$i-1][$j-1][0]-$prices[$i]);
            }
        }
        return $dp[count($prices)-1][$k][0];
    }
}
$ret = (new Solution())->maxProfit(
    2, [2,4,1]
//    2, [3,2,6,5,0,3]
);
var_dump($ret);
