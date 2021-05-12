<?php
//给你一个整数数组 nums 。
//如果一组数字 (i,j) 满足 nums[i] == nums[j] 且 i < j ，就可以认为这是一组 好数对 。
//返回好数对的数目。

class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function numIdenticalPairs($nums) {
        $map = array();
        foreach($nums as $k => $num){
            $map[$num]++;
        }
        $ret = 0;
        foreach($map as $count){
            $ret += $count * ($count-1) / 2;
        }
        return $ret;
    }
}

$result = (new Solution())->numIdenticalPairs([1,2,3,1,1,3]);
var_dump($result);
//[1,2,3,1,1,3] 解释：有 4 组好数对，分别是 (0,3), (0,4), (3,4), (2,5) ，下标从 0 开始
