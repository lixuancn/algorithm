<?php

class Solution {

    /**
     * @param Integer[][] $matrix
     * @param Integer $target
     * @return Boolean
     */
    function searchMatrix($matrix, $target) {
        $m = count($matrix);
        $n = count($matrix[0]);
        $lo = 0; $hi = $m * $n - 1;
        while($lo <= $hi){
            $mid = $lo + floor(($hi - $lo)/2);
            $k1 = intval($mid / $n);
            $k2 = intval($mid % $n);
            $midValue = $matrix[$k1][$k2];
            if($midValue == $target){
                return true;
            }
            if($target > $midValue){
                $lo = $mid + 1;
            }else{
                $hi = $mid - 1;
            }
        }
        return false;
    }
}

$ret = (new Solution())->searchMatrix(
    [[1], [3]],3
);
var_dump($ret);

