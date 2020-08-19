<?php
//查并集

class Solution {
    private $p = array();
    private $ret = 0;
    /**
     * @param Integer[][] $M
     * @return Integer
     */
    function findCircleNum($M) {
        if(!$M){
            return 0;
        }
        $n = count($M);
        for($i=0; $i<$n; $i++){
            $this->p[$i] = $i;
        }
        for($i=0; $i<$n; $i++) {
            for ($j = 0; $j < $n; $j++) {
                if($M[$i][$j] == 1){
                    $this->union($i, $j);
                }
            }
        }
        $ret = array();
        for($i=0; $i<$n; $i++) {
            $r = $this->parent($i);
            $ret[$r] = 1;
        }
        return count($ret);
    }

    private function union($i, $j){
        $p1 = $this->parent($i);
        $p2 = $this->parent($j);
        $this->p[$p2] = $p1;
    }

    private function parent($i){
        $root = $i;
        while($this->p[$root] != $root){
            $root = $this->p[$root];
        }
        while($this->p[$i] != $i){
            $x = $i;
            $i = $this->p[$i];
            $this->p[$x] = $root;
        }
        return $root;
    }
}

$obj = new Solution();
$r = $obj->findCircleNum(
    [
        [1,1,0],
        [1,1,0],
        [0,0,1]
    ]
);
var_dump($r);
