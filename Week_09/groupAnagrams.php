<?php
class Solution {
    function groupAnagrams($strs) {
        if(!$strs) return [];
        $ans = array();
        foreach($strs as $s){
            $map = array_fill(0, count($s), 0);
            for($i=0; $i<strlen($s); $i++){
                $k = ord($s[$i]) - ord('a');
                $map[$k] = isset($map[$k]) ? $map[$k]+1 : 1;
            }
            $key = '';
            for($i=0; $i<26; $i++){
                $key .= $map[$i] . '_';
            }
            $ans[$key][] = $s;
        }
        return array_values($ans);
    }
}
$obj = new Solution();
$r = $obj->groupAnagrams(
    ["ddddddddddg","dgggggggggg"]
);
var_dump($r);
