<?php

/**
 * 循环两遍nums
 * 不需要知道谁加谁等于target，而是用target-当前循环的值得到所需要的值。
 * 把nums转成map，这样在查找所需要的值时只用O(1)。
 * Class Solution
 */
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    function twoSum($nums, $target) {
        $map = array();
        foreach($nums as $i => $num){
            $map[$num] = $i;
        }
        foreach($nums as $i => $num){
            $complement = $target - $num;
            if(isset($map[$complement]) && $map[$complement] != $i){
                return array($i, $map[$complement]);
            }
        }
        return array();
    }
}
$obj = new Solution();
$nums = $obj->twoSum([2, 7, 11, 15], 9);
print_r($nums);
