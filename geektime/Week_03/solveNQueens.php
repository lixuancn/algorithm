<?php
//回溯
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
        $this->dfs($n, 0, []);
        $ret = [];
        foreach($this->ret as $ans){
            $one = array();
            foreach($ans as $r){
                $row = array_fill(0, $n, '.');
                $row[$r] = 'Q';
                $one[] = implode('', $row);
            }
            $ret[] = $one;
        }
        return $ret;
    }

    private function dfs($n, $row, $ans)
    {
        if ($row >= $n){
            $this->ret[] = $ans;
            return;
        }
        for($col=0; $col<$n; $col++){
            if(isset($this->col[$col]) || isset($this->pie[$row+$col]) || isset($this->na[$row-$col])){
                continue;
            }
            $this->col[$col] = 1;
            $this->pie[$row+$col] = 1;
            $this->na[$row-$col] = 1;
            $ans[] = $col;
            $this->dfs($n, $row+1, $ans);
            array_pop($ans);
            unset($this->col[$col]);
            unset($this->pie[$row+$col]);
            unset($this->na[$row-$col]);
        }
    }
}
$ret = (new Solution())->solveNQueens(4);
var_dump($ret);
