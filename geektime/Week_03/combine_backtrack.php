<?php
//回溯
class Solution {

    private $ret = array();
    /**
     * @param Integer $n
     * @param Integer $k
     * @return Integer[][]
     */
    function combine($n, $k) {
        $this->backtrack(1, $n, $k, array());
        return $this->ret;
    }

    function backtrack($i, $n, $k, $ans){
        if(count($ans) == $k){
            $this->ret[] = $ans;
            return;
        }
        for($j = $i; $j<=$n; $j++){
            $ans[] = $j;
            $this->backtrack($j+1, $n, $k, $ans);
            array_pop($ans);
        }
        return;
    }
}

$ret = (new Solution())->combine(4, 2);
var_dump($ret);
