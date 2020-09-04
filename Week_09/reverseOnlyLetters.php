<?php
//反转字符串中的单词 III
class Solution {
    /**
     * @param String $S
     * @return String
     */
    function reverseOnlyLetters($S) {
        if(!$S) return $S;
        $i = 0;
        $j = strlen($S) - 1;
        while($i < $j){
            if(!$this->isLetter($S[$i])){
                $i++;
            }
            if(!$this->isLetter($S[$j])){
                $j--;
            }
            if($this->isLetter($S[$i]) && $this->isLetter($S[$j])){
                $tmp = $S[$i];
                $S[$i] = $S[$j];
                $S[$j] = $tmp;
                $i++;
                $j--;
            }
        }
        return $S;
    }
    private function isLetter($c){
        $o = ord($c);
        if((ord('A') <= $o && $o <= ord('Z'))
            || (ord('a') <= $o && $o <= ord('z'))
        ){
            return true;
        }
        return false;
    }
}
$obj = new Solution();
$r = $obj->reverseOnlyLetters(
//    "ab-cd"
//    "a-bC-dEf-ghIj"
    "Test1ng-Leet=code-Q!"
);
var_dump($r);
