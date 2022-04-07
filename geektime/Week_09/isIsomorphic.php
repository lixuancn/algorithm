<?php
//思路：https://leetcode-cn.com/problems/regular-expression-matching/solution/ji-yu-guan-fang-ti-jie-gen-xiang-xi-de-jiang-jie-b/
class Solution {
    /**
     * @param String $s
     * @param String $t
     * @return Boolean
     */
    function isIsomorphic($s, $t) {
        if($s == $t) return true;
        if(strlen($s) != strlen($t)) return false;
        $m = strlen($s);
        $n = strlen($t);
        $map = array();
        $need = array();
        for($i=0; $i<$m; $i++){
            $map[$s[$i]][] = $i;
            $need[$t[$i]][] = $i;
        }
        $needArr = array_values($need);
        foreach(array_values($map) as $key => $posList){
            if(!isset($needArr[$key]) || count($needArr[$key]) != count($posList)){
                return false;
            }
            foreach($posList as $k=>$pos){
                if($pos != $needArr[$key][$k]){
                    return false;
                }
            }
        }
        return true;
    }
}
$obj = new Solution();
$r = $obj->isIsomorphic(
//    "egg", "add"
//    "foo", "bar"
    "paper", "title"
);
var_dump($r);
