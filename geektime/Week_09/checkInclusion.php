<?php
//滑动窗口 左闭右开
class Solution {
    /**
     * @param String $s1
     * @param String $s2
     * @return Boolean
     */
    function checkInclusion($s1, $s2) {
        $need = array();
        $m = strlen($s1);
        $n = strlen($s2);
        for($i=0; $i<$m; $i++) {
            $need[$s1[$i]]++;
        }
        //左开右闭区间
        $left = $right = 0;
        $window = array();
        $valid = 0;
        $start = 0;
        $len = null;
        while($right < $n) {
            $c = $s2[$right];
            $right++;
            if (isset($need[$c])) {
                $window[$c]++;
                if ($window[$c] == $need[$c]) {
                    $valid++;
                }
            }
            while($right - $left >= $m){
                if ($valid == count($need)) {
                    return true;
                }
                $d = $s2[$left];
                $left++;
                if(isset($need[$d])){
                    if($window[$d] == $need[$d]){
                        $valid--;
                    }
                    $window[$d]--;
                }
            }
        }
        return false;
    }
}
$obj = new Solution();
$r = $obj->checkInclusion(
    "ab", "eidboaoo"
//    "ab", "eidbaooo"
//    "hello", "ooolleoooleh"
);
var_dump($r);
