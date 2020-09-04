<?php
//倒序计数法
class Solution {
    /**
     * @param String $s
     * @return Integer
     */
    function lengthOfLastWord($s) {
        if(!$s) return 0;
        $n = strlen($s);
        $len = 0;
        for($i=$n-1; $i>=0; $i--){
            if(($len == 0 && $s[$i] != ' ') || ($len > 0 && $s[$i] != ' ')){
                $len++;
            }
            if($len > 0 && $s[$i] == ' '){
                break;
            }
        }
        return $len;
    }
}
$obj = new Solution();
$r = $obj->lengthOfLastWord(
//    "Hello World"
    "b   a    "
);
var_dump($r);
