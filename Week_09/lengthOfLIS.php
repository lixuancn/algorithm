<?php
//DP
class Solution {
    public function lengthOfLIS($nums){
        $n = count($nums);
        $dp = array_fill(0, $n, 1);
        $max = 0;
        for($i=0; $i<$n; $i++){
            for($j=0; $j<$i; $j++){
                if($nums[$i] > $nums[$j]){
                    $dp[$i] = max($dp[$i], $dp[$j]+1);
                }
            }
            $max = max($dp[$i], $max);
        }
        return $max;
    }
}
$obj = new Solution();
$r = $obj->lengthOfLIS(
//    [-2, -1]
    [10,9,2,5,3,7,101,18]
);
var_dump($r);
