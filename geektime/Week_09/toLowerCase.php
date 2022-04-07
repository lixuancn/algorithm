<?php
class Solution {
    /**
     * @param String $str
     * @return String
     */
    function toLowerCase($str) {
        if(!$str) return '';
        $n = strlen($str);
        for($i=0; $i<$n; $i++){
            if(ord($str[$i]) >= ord('A') && ord($str[$i]) <= ord('Z')){
                $str[$i] = chr(ord($str[$i]) + 32);
            }
        }
        return $str;
    }
}
$obj = new Solution();
$r = $obj->toLowerCase(
    "Hello"
);
var_dump($r);
