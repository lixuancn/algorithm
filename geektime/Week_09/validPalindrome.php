<?php
class Solution {
    /**
     * @param String $s
     * @return Boolean
     */
    function validPalindrome($s) {
        if(!$s) return true;
        $left = 0;
        $right = strlen($s) - 1;
        while($left < $right){
            if($s[$left] == $s[$right]) {
                $left++;
                $right--;
                continue;
            }
            $flag1 = $flag2 = true;
            for($i=$left,$j=$right-1; $i<$j; $i++,$j--){
                if($s[$i] != $s[$j]){
                    $flag1 = false;
                    break;
                }
            }
            for($i=$left+1,$j=$right; $i<$j; $i++,$j--){
                if($s[$i] != $s[$j]){
                    $flag2 = false;
                    break;
                }
            }
            return $flag1 || $flag2;
        }
        return true;
    }
}
$obj = new Solution();
$r = $obj->validPalindrome(
//    'aba'
//    'abca'
//    "deeee"
    'abc'
);
var_dump($r);
