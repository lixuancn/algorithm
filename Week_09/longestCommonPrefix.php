<?php

class Solution {
    /**
     * @param String[] $strs
     * @return String
     */
    function longestCommonPrefix($strs) {
        if(!$strs) return '';
        $n = strlen($strs[0]);
        $prefix = '';
        for($i=0; $i<$n; $i++){
            $tmp = $strs[0][$i];
            for($j=1; $j<count($strs); $j++){
                if(isset($strs[$j][$i]) && $strs[$j][$i] != $tmp){
                    break 2;
                }
            }
            $prefix .= $tmp;
        }
        return $prefix;
    }
}
$obj = new Solution();
$r = $obj->longestCommonPrefix(
    ["flower","flow","flight"]
);
var_dump($r);
