<?php
class Solution {
    /**
     * @param String $J
     * @param String $S
     * @return Integer
     */
    function numJewelsInStones($J, $S) {
        $map = array();
        for($i=0; $i<strlen($J); $i++){
            $map[$J[$i]] = 1;
        }
        $count = 0;
        for($i=0; $i<strlen($S); $i++){
            if(isset($map[$S[$i]])){
                $count++;
            }
        }
        return $count;
    }
}
$obj = new Solution();
$r = $obj->numJewelsInStones(
    "aA", "aAAbbbb"
);
var_dump($r);
