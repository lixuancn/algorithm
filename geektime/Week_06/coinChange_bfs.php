<?php
//bfs
class Solution {

    private $min = 0;
    /**
     * @param Integer[] $coins
     * @param Integer $amount
     * @return Integer
     */
    function coinChange($coins, $amount)
    {
        $queue = array();
        array_push($queue, array($amount, 0));
        while($queue){
            list($amount, $level) = array_shift($queue);
            if($amount == 0){
                return $level;
            }
            if($amount < 0){
                return -1;
            }
            foreach($coins as $coin){
                array_push($queue, array($amount-$coin, $level+1));
            }
        }
        return -1;
    }
}
$ret = (new Solution())->coinChange(
//    [1, 2, 5], 11
    [186,419,83,408], 6249
//    [2], 3
);
var_dump($ret);
