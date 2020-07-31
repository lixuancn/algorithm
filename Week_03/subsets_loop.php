<?php
//迭代
class Solution {
    private $ret = array();
    /**
     * @param Integer[] $nums
     * @return Integer[][]
     */
    function subsets($nums) {
        $this->ret[] = [];
        foreach($nums as $num){
            $newSubsets = array();
            foreach($this->ret as $curr){
                array_push($curr, $num);
                $newSubsets[] = $curr;
            }
            foreach($newSubsets as $sub){
                $this->ret[] = $sub;
            }
        }
        return $this->ret;
    }
}

$ret = (new Solution())->subsets([1,2,3]);
var_dump($ret);
