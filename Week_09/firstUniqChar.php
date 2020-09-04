<?php

class Solution {
    /**
     * @param String $s
     * @return Integer
     */
    function firstUniqChar($s) {
        $countMap = array();
        for($i=0; $i<strlen($s); $i++){
            $countMap[$s[$i]]++;
        }
        for($i=0; $i<strlen($s); $i++){
            if($countMap[$s[$i]] == 1){
                return $i;
            }
        }
        return -1;
    }
}
$obj = new Solution();
$r = $obj->firstUniqChar(
    'loveleetcode'
);
var_dump($r);
