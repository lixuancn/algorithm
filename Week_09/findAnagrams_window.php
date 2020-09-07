<?php
//滑动窗口 左闭右开
class Solution {
    /**
     * @param String $s
     * @param String $p
     * @return Integer[]
     */
    function findAnagrams($s, $p) {
        $ans = array();
        $need = array();
        $m = strlen($s);
        $n = strlen($p);
        for($i=0; $i<$n; $i++) {
            $need[$p[$i]]++;
        }
        //左开右闭区间
        $left = $right = 0;
        $window = array();
        $valid = 0;
        $start = 0;
        $len = null;
        while($right < $m) {
            $c = $s[$right];
            $right++;
            if (isset($need[$c])) {
                $window[$c]++;
                if ($window[$c] == $need[$c]) {
                    $valid++;
                }
            }
            while($right - $left >= $n){
                if ($valid == count($need)) {
                    $ans[] = $left;
                }
                $d = $s[$left];
                $left++;
                if(isset($need[$d])){
                    if($window[$d] == $need[$d]){
                        $valid--;
                    }
                    $window[$d]--;
                }
            }
        }
        return $ans;
    }
}
$obj = new Solution();
$r = $obj->findAnagrams(
    "cbaebabacd", "abc"
);
var_dump($r);
