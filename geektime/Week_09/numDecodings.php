<?php
//DP
class Solution {
    /**
     * @param String $s
     * @return Integer
     */
    function numDecodings($s) {
        if(!strlen($s) || $s[0] == 0){
            return 0;
        }
        $pre = $cur = 1;
        for($i=1; $i<strlen($s); $i++){
            $tmp = $cur;
            if($s[$i] == 0){
                if($s[$i-1] != 1 && $s[$i-1] != 2){
                    return 0;
                }
                $cur = $pre;
            }else if($s[$i-1] == 1 || ($s[$i-1] == 2 && $s[$i] <= 6)){
                $cur = $pre + $cur;
            }
            $pre = $tmp;
        }
        return $cur;
    }
}
$obj = new Solution();
$r = $obj->numDecodings(
    '226'
);
var_dump($r);
