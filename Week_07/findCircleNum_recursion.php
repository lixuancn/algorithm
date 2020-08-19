<?php
//递归

class Solution {
    private $visited = array();
    private $ret = 0;
    /**
     * @param Integer[][] $M
     * @return Integer
     */
    function findCircleNum($M) {
        if(!$M){
            return 0;
        }
        for($i=0; $i<count($M); $i++){
            if(!isset($this->visited[$i])){
                $this->dfs($M, $i);
                $this->ret++;
            }
        }
        return $this->ret;
    }

    function dfs($M, $i){
        for($j=0; $j<count($M); $j++){
            if($M[$i][$j] == 1 && !isset($this->visited[$j])){
                $this->visited[$j] = 1;
                if($i==$j){
                    continue;
                }
                $this->dfs($M, $j);
            }
        }
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
