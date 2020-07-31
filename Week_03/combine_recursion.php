<?php
//递归
class Solution {

    private $ret = array();
    /**
     * @param Integer $n
     * @param Integer $k
     * @return Integer[][]
     */
    function combine($n, $k) {
        for($i=1; $i<=$n; $i++){
            $this->helper($i, $n, $k, array());
        }
        var_dump($this->ret);
        return $this->ret;
    }

    function helper($i, $n, $k, $ans){
        $ans[] = $i;
        if(count($ans) == $k){
            $this->ret[] = $ans;
            return;
        }
        for($j = $i + 1; $j<=$n; $j++){
            $this->helper($j, $n, $k, $ans);
        }
        return;
    }
}
