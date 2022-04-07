<?php
//递归
class Solution {
    private $ret = array();

    /**
     * @param Integer $m
     * @param Integer $n
     * @return Integer
     */
    function uniquePaths($m, $n) {
        $this->helper($m, $n, 0, 0);
        return $this->ret[0][0];
    }

    private function helper($m, $n, $i, $j){
        if($i >= $m-1 || $j >= $n -1){
            $this->ret[$i][$j] = 1;
            return 1;
        }
        if(isset($this->ret[$i][$j])){
            return $this->ret[$i][$j];
        }
        $this->ret[$i+1][$j] = $this->helper($m, $n, $i+1, $j);
        $this->ret[$i][$j+1] = $this->helper($m, $n, $i, $j+1);
        $this->ret[$i][$j] = $this->ret[$i+1][$j] + $this->ret[$i][$j+1];
        return $this->ret[$i][$j];
    }
}
$ret = (new Solution())->uniquePaths(
    23,12
);
var_dump($ret);

