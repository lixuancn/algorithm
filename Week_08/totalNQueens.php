<?php
class Solution {
    private $ret = 0;
    /**
     * @param Integer $n
     * @return Integer
     */
    function totalNQueens($n) {
        if($n < 0){
            return [];
        }
        $this->dfs($n, 0, array());
        return $this->ret;
    }

    function dfs($n, $i, $ans){
        if($i>=$n){
            $this->ret++;
            return;
        }
        for($j=0; $j<$n; $j++){
            if(isset($this->col[$j]) || isset($this->pie[$i+$j]) || isset($this->na[$i-$j])){
                continue;
            }
            $this->col[$j] = 1;
            $this->pie[$i+$j] = 1;
            $this->na[$i-$j] = 1;
            $ans[] = $j;
            $this->dfs($n, $i+1, $ans);
            array_pop($ans);
            unset($this->col[$j]);
            unset($this->pie[$i+$j]);
            unset($this->na[$i-$j]);
        }
    }
}
