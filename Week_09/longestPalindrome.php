<?php
//DP
class Solution {
    /**
     * @param String $s
     * @return String
     */
    function longestPalindrome($s) {
        if(!$s) return $s;
        $ans = '';
        $dp = array();
        for($i=strlen($s)-1; $i>=0; $i--){
            for($j=$i; $j<strlen($s); $j++){
                //i~j的子串是不是回文串
                $dp[$i][$j] = $s[$i] == $s[$j] && ($j - $i < 2 || $dp[$i+1][$j-1]);
                if($dp[$i][$j] && $j-$i+1 > strlen($ans)){
                    $ans = substr($s, $i, $j-$i+1);
                }
            }
        }
        return $ans;
    }
}
$obj = new Solution();
$r = $obj->longestPalindrome(
    'babad'
//    'cbbd'
);
var_dump($r);
