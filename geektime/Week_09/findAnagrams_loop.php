<?php
//暴力 leetcode上超时
class Solution {
    /**
     * @param String $s
     * @param String $p
     * @return Integer[]
     */
    function findAnagrams($s, $p) {
        $ans = array();
        $map = array();
        $m = strlen($s);
        $n = strlen($p);
        for($j=0; $j<$n; $j++) {
            $map[$p[$j]]++;
        }
        for($i=0; $i<$m-$n+1; $i++){
            $tmp = $map;
            for($j=$i; $j<$i+$n; $j++) {
                if(!isset($tmp[$s[$j]])){
                    break;
                }
                $tmp[$s[$j]]--;
                if($tmp[$s[$j]] < 0){
                    break;
                }
            }
            $result = true;
            foreach($tmp as $t){
                if($t != 0){
                    $result = false;
                    break;
                }
            }
            if($result){
                $ans[] = $i;
            }
        }
        return $ans;
    }
}
$obj = new Solution();
$r = $obj->findAnagrams(
    "cbaebabacd", "abc"
//    "abab", "ab"
);
var_dump($r);
