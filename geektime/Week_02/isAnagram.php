<?php
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
        //leetcode 有bug，这里应该还要判断两个字符串不相等才叫异位词
//        if($s == $t){
//            return false;
//        }
        $map = array();
        for($i=0; $i<strlen($s); $i++){
            if(!isset($map[$s[$i]])){
                $map[$s[$i]] = 0;
            }
            if(!isset($map[$t[$i]])){
                $map[$t[$i]] = 0;
            }
            $map[$s[$i]]++;
            $map[$t[$i]]--;
        }
        foreach($map as $m){
            if($m != 0){
                return false;
            }
        }
        return true;
    }
}

$r = (new Solution())->isAnagram("anagram", "nagaram");
var_dump($r);
