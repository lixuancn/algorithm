<?php

/**
 * 循环一遍nums，再循环两遍nums的基础上优化。
 * 在nums转map时就查看所需要的值是否已经在map中了。
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
            $complement = $target - $num;
            if(isset($map[$complement])){
                return array($i, $map[$complement]);
            }
            $map[$num] = $i;
        }
        return array();
    }
}
$obj = new Solution();
$nums = $obj->twoSum([2, 7, 11, 15], 9);
print_r($nums);
