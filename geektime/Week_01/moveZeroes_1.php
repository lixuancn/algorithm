<?php

/**
 * 新数组来记录非0，然后空位用0填充，最后复制回$nums
 * Class Solution
 */
class Solution {

    /**
     * @param Integer[] $nums
     * @return NULL
     */
    function moveZeroes(&$nums) {
        $zeroNum = 0;
        foreach($nums as $num){
            if($num == 0){
                $zeroNum++;
            }
        }
        $newArray = array();
        foreach($nums as $num){
            if($num != 0){
                $newArray[] = $num;
            }
        }
        for($i=0; $i<$zeroNum; $i++){
            $newArray[] = 0;
        }
        $nums = $newArray;
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
$nums = [0,1,0,3,12];
$nums = [0,0,1];
$obj->moveZeroes($nums);
print_r($nums);
