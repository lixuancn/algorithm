<?php
//贪心
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function jump($nums) {
        $ans = 0;
        $maxPos = 0;
        $end = 0;
        for($i=0; $i<count($nums)-1; $i++){
            $maxPos = max($maxPos, $nums[$i]+$i);
            if($i == $end){
                $end = $maxPos;
                $ans++;
            }
        }
        return $ans;
    }
}

$ret = (new Solution())->jump(
    [2,3,1,1,4]
);
var_dump($ret);
