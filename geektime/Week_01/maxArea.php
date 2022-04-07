<?php
class Solution {

    /**
     * @param Integer[] $height
     * @return Integer
     */
    function maxArea($height) {
        $max = 0;
        $left = 0;
        $right = count($height) - 1;
        while ($left < $right){
            $area = ($height[$left] > $height[$right] ? $height[$right] : $height[$left]) * ($right - $left);
            if($max < $area){
                $max = $area;
            }
            if($height[$left] >= $height[$right]){
                $right --;
            }else{
                $left ++;
            }
        }
        return $max;
    }
}
$obj = new Solution();
echo $obj->maxArea([1,8,6,2,5,4,8,3,7]);
