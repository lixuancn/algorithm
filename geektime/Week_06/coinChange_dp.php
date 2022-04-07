<?php
//dp
class Solution {

    /**
     * @param Integer[] $coins
     * @param Integer $amount
     * @return Integer
     */
    function coinChange($coins, $amount) {
        if(!$amount){
            return 0;
        }
        $max = $amount + 1;
        $dp = array_fill(0, $amount+1, $max);
        $dp[0] = 0;
        //f(n) = min(f(n-k), k range coins) + 1
        for($i=1; $i<=$amount; $i++){
            foreach($coins as $coin){
                if($coin <= $i){
                    $dp[$i] = min($dp[$i], $dp[$i-$coin]+1);
                }
            }
        }
        return $dp[$amount] > $amount ? -1 : $dp[$amount];
    }
}
$ret = (new Solution())->coinChange(
//    [1, 2, 5], 11
//    [186,419,83,408], 6249
    [2], 3
);
var_dump($ret);
