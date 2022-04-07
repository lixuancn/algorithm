<?php
//位运算解N皇后
class Solution {
    private $ret = 0;
    /**
     * @param Integer $n
     * @return Integer
     */
    function totalNQueens($n) {
        if($n < 0){
            return 0;
        }
        $this->dfs($n, 0, 0, 0, 0);
        return $this->ret;
    }

    //$cols, $pie, $na 表示列撇捺上被占用的点，占用的是1
    function dfs($n, $i, $cols, $pie, $na){
        if($i>=$n){
            $this->ret++;
            return;
        }
        //得到当前所有的空位
        //$cols|$pie|$na表示所有被占用的，~($cols|$pie|$na)所有可用的
        //(1<<$n)-1
        $bits = (~($cols|$pie|$na)) & ((1<<$n)-1);
        while($bits){
            $p = $bits & -$bits; # 取到最低位的1
            $bits = $bits & ($bits - 1); # 表示在p位置上放入皇后
            $this->dfs($n, $i+1, $cols | $p, ($pie | $p) << 1, ($na | $p) >> 1);
        }
    }
}

$obj = new Solution();
$r = $obj->totalNQueens(
    4
);
var_dump($r);
