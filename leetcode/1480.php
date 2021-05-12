<?php
//给你一个数组 nums 。数组「动态和」的计算公式为：runningSum[i] = sum(nums[0]…nums[i]) 。
//请返回 nums 的动态和。

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer[]
     */
    function runningSum($nums) {
        foreach($nums as $k => $num){
            $nums[$k] += $nums[$k-1];
        }
        return $nums;
    }
}

$result = (new Solution())->runningSum([1,2,3,4]);
var_dump($result);
