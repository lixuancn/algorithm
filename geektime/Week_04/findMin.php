<?php

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function findMin($nums) {
        if(count($nums) == 1){
            return $nums[0];
        }
        if(count($nums) == 2){
            return min($nums[0], $nums[1]);
        }
        $lo = 0; $hi = count($nums) - 1;
        if($nums[0] < $nums[$hi]){
            return $nums[0];
        }
        while($lo <= $hi){
            $mid = ceil($lo + ($hi-$lo)/2);
            if($nums[$mid+1] && $nums[$mid] > $nums[$mid+1]){
                return $nums[$mid+1];
            }
            if(isset($nums[$mid-1]) && $nums[$mid] < $nums[$mid-1]){
                return $nums[$mid];
            }
            //左边是有序的，就去右边
            if($nums[0] < $nums[$mid]){
                $lo = $mid + 1;
            }else{
                $hi = $mid - 1;
            }
        }
        return $nums[$lo];
    }
}

$ret = (new Solution())->findMin(
    [9,0,2,7,8]
);
var_dump($ret);

