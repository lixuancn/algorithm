<?php
//滑动窗口 左闭右开
class Solution {
    /**
     * @param String $s
     * @param String $p
     * @return Integer[]
     */
    function lengthOfLongestSubstring($s) {
        $ans = 0;
        $m = strlen($s);
        //左开右闭区间
        $left = $right = 0;
        $window = array();
        while($right < $m) {
            $c = $s[$right];
            $right++;
            $window[$c]++;
            while($window[$c] > 1){
                $d = $s[$left];
                $left++;
                $window[$d]--;
            }
            $ans = max($ans, $right-$left);
        }
        return $ans;
    }
}
$obj = new Solution();
$r = $obj->lengthOfLongestSubstring(
    "abcabcbb"
);
var_dump($r);
