<?php
//贪心
class Solution {

    /**
     * @param Integer[] $nums
     * @return Boolean
     */
    function canJump($nums) {
        if(!$nums) return false;
        $endPos = count($nums) - 1;
        for($i=count($nums)-2; $i>=0; $i--){
            if($nums[$i] + $i >= $endPos){
                $endPos = $i;
            }
        }
        return $endPos == 0;
    }
}

$ret = (new Solution())->canJump(
    [2,3,1,1,4]
);
var_dump($ret);
