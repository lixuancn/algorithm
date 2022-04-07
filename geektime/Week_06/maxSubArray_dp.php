<?php
//dp
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function maxSubArray($nums) {
        //f(i) = f(i) + max(f(i-1), 0)
        $max = $nums[0];
        for($i=1; $i<count($nums); $i++){
            $nums[$i] = $nums[$i] + max($nums[$i-1], 0);
            if($nums[$i] > $max){
                $max = $nums[$i];
            }
        }
        return $max;
    }
}
$ret = (new Solution())->maxSubArray(
    [-2,1,-3,4,-1,2,1,-5,4]
);
var_dump($ret);
