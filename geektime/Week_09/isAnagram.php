<?php
//异位词
class Solution {
    /**
     * @param String $s
     * @param String $t
     * @return Boolean
     */
    function isAnagram($s, $t) {
        if(strlen($s) != strlen($t)){
            return false;
        }
        $map = array();
        for($i=0; $i<strlen($s); $i++){
            $map[$s[$i]] = isset($map[$s[$i]]) ? $map[$s[$i]]+1 : 1;
            $map[$t[$i]] = isset($map[$t[$i]]) ? $map[$t[$i]]-1 : -1;
        }
        foreach($map as $m){
            if($m != 0){
                return false;
            }
        }
        return true;
    }
}
$obj = new Solution();
$r = $obj->isAnagram(
    'anagram', 'nagaram'
);
var_dump($r);
