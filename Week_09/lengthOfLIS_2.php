<?php
//贪心 + 二分查找
class Solution {
    public function lengthOfLIS($nums){
        $d = array();
        foreach($nums as $n){
            if(!$d || $n > $d[count($d)-1]){
                $d[] = $n;
            }else{
                $l = 0;
                $r = count($d) - 1;
                $loc = $r;
                while($l <= $r){
                    $mid = intval(($l + $r) / 2);
                    if($d[$mid] >= $n){
                        $loc = $mid;
                        $r = $mid - 1;
                    }else{
                        $l = $mid + 1;
                    }
                }
                $d[$loc] = $n;
            }
        }
        return count($d);
    }
}
$obj = new Solution();
$r = $obj->lengthOfLIS(
//    [-2, -1]
    [10,9,2,5,3,7,101,18]
);
var_dump($r);
