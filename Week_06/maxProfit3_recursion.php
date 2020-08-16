<?php
//递归
class Solution {
    private $cache = array();
    /**
     * @param Integer[] $prices
     * @return Integer
     */
    function maxProfit($prices) {
        if(!$prices || count($prices)<2){
            return 0;
        }
        return $this->dfs($prices, 0, 0, 0);
    }

    private function dfs($prices, $index, $status, $k){
        if($index == count($prices) || $k == 2){
            return 0;
        }
        if(isset($this->cache["{$index}_{$status}_{$k}"])){
            return $this->cache["{$index}_{$status}_{$k}"];
        }
        $a = $b = $c = 0;
        //保持不动
        $a = $this->dfs($prices, $index+1, $status, $k);
        //卖
        if($status == 1){
            $b = $this->dfs($prices, $index+1, 0, $k+1) + $prices[$index];
        }else{
            //买
            $c = $this->dfs($prices, $index+1, 1, $k) - $prices[$index];
        }
        $r = max($a, $b, $c);
        $this->cache["{$index}_{$status}_{$k}"] = $r;
        return $r;
    }
}
$ret = (new Solution())->maxProfit(
    [3,3,5,0,0,3,1,4]
);
var_dump($ret);
