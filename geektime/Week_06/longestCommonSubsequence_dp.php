<?php
//DP
class Solution {
    /**
     * @param String $text1
     * @param String $text2
     * @return Integer
     */
    function longestCommonSubsequence($text1, $text2) {
        $m = strlen($text1);
        $n = strlen($text2);
        if(!$m || !$n){
            return 0;
        }
        $ret = array_fill(0, $m, array_fill(0, $n, 0));
        for($i=0; $i<$m; $i++){
            for($j=0; $j<$n; $j++){
                if($text1[$i] == $text2[$j]){
                    $ret[$i][$j] = 1 + intval($ret[$i-1][$j-1]);
                }else{
                    $ret[$i][$j] = max(intval($ret[$i-1][$j]), intval($ret[$i][$j-1]));
                }
            }
        }
        return $ret[$m-1][$n-1];
    }
}
$ret = (new Solution())->longestCommonSubsequence(
    "pmjghexybyrgzczy", "hafcdqbgncrcbihkd"
//    "abcde", "ace"
//    "abc", "def"
);
var_dump($ret);

