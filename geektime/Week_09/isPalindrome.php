<?php
class Solution {
    /**
     * @param String $s
     * @return Boolean
     */
    function isPalindrome($s) {
        if(!$s) return true;
        $str = '';
        for($i=0; $i<strlen($s); $i++){
            $ord = ord($s[$i]);
            if(($ord>=ord('a') && $ord<=ord('z')) ||
                $ord>=ord('A') && $ord<=ord('Z') ||
                $ord>=ord('0') && $ord<=ord('9')
            ){
                $str .= strtolower($s[$i]);
            }
        }
        $mid = intval(strlen($str)/2);
        if(strlen($str) % 2 == 0){
            $left = $mid - 1;
            $right = $mid;
        }else{
            $left = $mid - 1;
            $right = $mid + 1;
        }
        while($left >= 0 && $right < strlen($s)){
            if($str[$left] != $str[$right]){
                return false;
            }
            $left--; $right++;
        }
        return true;
    }
}
$obj = new Solution();
$r = $obj->isPalindrome(
//    "A man, a plan, a canal: Panama"
    "race a car"
);
var_dump($r);
