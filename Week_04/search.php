<?php

class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer
     */
    function search($nums, $target) {
        $lo = 0;
        $hi = count($nums) - 1;
        while($lo < $hi){
            $mid = intval($lo + ($hi - $lo)/2);
            //左边是有序的，但是需要去右边找的情况。
            if($nums[0] <= $nums[$mid] && ($target > $nums[$mid] || $target < $nums[0])){
                $lo = $mid + 1;
            }else if($target > $nums[$mid] && $target < $nums[0]){
                //左边是无序的，需要在右边找
                $lo = $mid + 1;
            }else{
                //左边是有序的，就在左边找
                //左边是无序的，就在左边找
                $hi = $mid;
            }
        }
        return $lo == $hi && $nums[$lo] == $target ? $lo : -1;
    }
}

$ret = (new Solution())->search(
    [4,5,6,7,0,1,2], 0
);
var_dump($ret);

