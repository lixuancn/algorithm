<?php

/**
 * 双指针，zero指向0的位置，cur是当前遍历的指针，zero每次替换时++，cur每次循环++
 * cur指向非0时，和zero替换。
 * 比方法一优化了空间复杂度，O(n)->O(1)
 * Class Solution
 */
class Solution {

    /**
     * @param Integer[] $nums
     * @return NULL
     */
    function moveZeroes(&$nums) {
        //$zero指向0的位置，cur是当前遍历的指针，cur指向非0时，和zero替换
        $cur = $zero = 0;
        for($cur=0, $noZero=0; $cur < count($nums); $cur++){
            if($nums[$cur] != 0){
                $this->swap($nums[$zero++], $nums[$cur]);
            }
        }
    }

    private function swap(&$i, &$j){
        $t = $i;
        $i = $j;
        $j = $t;
        return array($i, $j);
    }
}
$obj = new Solution();
$nums = [0,1,0,3,12];
$nums = [0,0,1];
$obj->moveZeroes($nums);
print_r($nums);
