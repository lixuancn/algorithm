<?php

/**
 * 遍历$nums，非0的就往前面放，那后面的全部都补0
 * 比方法一优化了空间复杂度，O(n)->O(1)
 * Class Solution
 */
class Solution {

    /**
     * @param Integer[] $nums
     * @return NULL
     */
    function moveZeroes(&$nums) {
        $nonZeroNum = 0;
        foreach($nums as $num){
            if($num != 0){
                $nums[$nonZeroNum++] = $num;
            }
        }
        for($i=$nonZeroNum; $i<count($nums); $i++){
            $nums[$i] = 0;
        }
        return;
    }

    private function swap(&$i, &$j){
        $t = $i;
        $i = $j;
        $j = $t;
        return array($i, $j);
    }
}
$obj = new Solution();
$nums = [0,0,1];
$nums = [0,1,0,3,12];
$obj->moveZeroes($nums);
print_r($nums);
