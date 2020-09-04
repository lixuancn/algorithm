<?php
class Solution {
    /**
     * @param String $str
     * @return Integer
     */
    function myAtoi($str) {
        if(!$str) return 0;
        $n = strlen($str);
        $head = 0;
        while($str[$head]==' ' && $head<$n){
            $head++;
        }
        $num = 0;
        $flag = 1;
        if($str[$head] == '+' || $str[$head] == '-'){
            $flag = $str[$head] == '-' ? -1 : 1;
            $head++;
        }
        $max = pow(2, 31) - 1;
        $min = -pow(2, 31);
        for($i=$head; $i<$n; $i++){
            if(in_array($str[$i], array('0','1','2','3','4','5','6','7','8','9'))){
                $digit = intval($str[$i]);
                if($max/10<$num || intval($max/10)==$num && $max%10<$digit){
                    return $flag > 0 ? $max : $min;
                }
                $num = $num * 10 + $digit;
            }else{
                break;
            }
        }
        return $num * $flag;
    }
}
$obj = new Solution();
$r = $obj->myAtoi(
//    '-42a'
//    "91283472332"
    "2147483648"
);
var_dump($r);
