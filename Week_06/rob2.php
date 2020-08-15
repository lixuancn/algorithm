<?php
//打家劫舍II，dp
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function rob($nums) {
        if(!$nums){
            return 0;
        }
        if(count($nums) < 2){
            return $nums[0];
        }
        $first = $nums[0];
        $second = max($first, $nums[1]);
        if(count($nums) < 3){
            return $second;
        }
        //结尾不偷
        for($i=2; $i<count($nums)-1; $i++){
            $cur = max($first+$nums[$i], $second);
            $first = $second;
            $second = $cur;
        }
        $p1 = $second;
        //开头不偷
        $first = $nums[1];
        $second = max($first, $nums[2]);
        for($i=3; $i<count($nums); $i++){
            $cur = max($first+$nums[$i], $second);
            $first = $second;
            $second = $cur;
        }
        $p2 = $second;
        return max($p1, $p2);
    }
}
$ret = (new Solution())->rob(
    [2,3,2]
//    [1,2,3,1]
);
var_dump($ret);
