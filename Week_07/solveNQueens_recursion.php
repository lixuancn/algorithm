<?php
//递归，剪枝
class Solution {
    private $ret = array();
    private $col = array();
    private $pie = array();
    private $na = array();
    /**
     * @param Integer $n
     * @return String[][]
     */
    function solveNQueens($n) {
        if($n < 0){
            return [];
        }
        $this->dfs($n, 0, array());
        $ret = array();
        foreach($this->ret as $ans){
            $r = array_fill(0, $n, array_fill(0, $n, '.'));
            foreach($ans as $i=>$j){
                $r[$i][$j] = 'Q';
                $r[$i] = implode('', $r[$i]);
            }
            $ret[] = $r;
        }
        return $ret;
    }

    function dfs($n, $i, $ans){
        if($i>=$n){
            $this->ret[] = $ans;
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
$obj = new Solution();
$r = $obj->solveNQueens(
    4
);
var_dump($r);
