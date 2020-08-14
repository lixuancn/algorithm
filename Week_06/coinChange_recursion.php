<?php
//递归
class Solution {

    private $min = 0;
    /**
     * @param Integer[] $coins
     * @param Integer $amount
     * @return Integer
     */
    function coinChange($coins, $amount)
    {
        $this->min = pow(2, 31);
        $this->helper($coins, $amount, 0);
        return $this->min;
    }

    function helper($coins, $amount, $count){
        if($amount == 0 && $this->min > $count){
            $this->min = $count;
            return;
        }
        if($amount < 0){
            return;
        }
        foreach($coins as $coin){
            $this->helper($coins, $amount-$coin, $count+1);
        }
    }
}
$ret = (new Solution())->coinChange(
    [1, 2, 5], 11
//    [186,419,83,408], 6249
//    [2], 3
);
var_dump($ret);
