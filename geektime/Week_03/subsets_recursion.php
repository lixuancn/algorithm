<?php
//递归
class Solution {
    private $ret = array();
    /**
     * @param Integer[] $nums
     * @return Integer[][]
     */
    function subsets($nums) {
        $this->dfs($nums, array(), 0);
        return $this->ret;
    }

    private function dfs($nums, $ans, $index){
        if($index == count($nums)){
            $this->ret[] = $ans;
            return;
        }
        $this->dfs($nums, $ans, $index+1);
        $ans[] = $nums[$index];
        $this->dfs($nums, $ans, $index+1);
    }
}

$ret = (new Solution())->subsets([1,2,3]);
var_dump($ret);
