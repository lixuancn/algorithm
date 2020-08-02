<?php
//回溯
//倒着走
class Solution {

    /**
     * @param Integer[] $nums
     * @return Boolean
     */
    function canJump($nums) {
        if(!$nums) return false;
        return $this->backtrack($nums, count($nums) - 1);
    }

    private function backtrack($nums, $target){
        if($target == 0){
            return true;
        }
        for($src=$target-1; $src>=0; $src--){
            if($nums[$src] >= $target - $src){
                return $this->backtrack($nums, $src);
            }
        }
        return false;
    }
}

$ret = (new Solution())->canJump(
    [2,3,1,1,4]
);
var_dump($ret);
