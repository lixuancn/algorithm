<?php
//暴力法
//但我的思路其实是反向贪心
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function jump($nums) {
        $ans = 0;
        for($i=count($nums)-1; $i>0; $i){
            $next = null;
            for($j=$i-1; $j>=0; $j--) {
                if($nums[$j] >= $i - $j){
                    $next = $j;
                }
            }
            if(is_null($next)){
                return 0;
            }
            $i = $next;
            $ans++;
        }
        return $ans;
    }
}

$ret = (new Solution())->jump(
    [2,3,1,1,4]
);
var_dump($ret);
