<?php
//暴力
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function maxSubArray($nums) {
        $max = 0;
        for($i=0; $i<count($nums); $i++){
            $sum = 0;
            for($j=$i; $j<count($nums); $j++){
                $sum += $nums[$j];
                if($sum>$max){
                    $max = $sum;
                }
            }
        }
        return $max;
    }
}
$ret = (new Solution())->maxSubArray(
    [-2,1,-3,4,-1,2,1,-5,4]
);
var_dump($ret);
