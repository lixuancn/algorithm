<?php
//翻转字符串里的单词
class Solution {
    /**
     * @param String $s
     * @return String
     */
    function reverseWords($s) {
        if(!$s) return $s;
        $n = strlen($s);
        $start = 0; $end = $n-1;
        while($start < $end && $s[$start] == ' ') $start++;
        while($start < $end && $s[$end] == ' ') $end--;
        $word = '';
        $wordList = array();
        while($start <= $end){
            if($s[$start] == ' '){
                if($word){
                    $wordList[] = $word;
                    $word = '';
                }
            }else{
                $word .= $s[$start];
                if($start == $end) $wordList[] = $word;
            }
            $start++;
        }
        $newStr = '';
        for($i=count($wordList)-1; $i>=0; $i--){
            $newStr .= $wordList[$i];
            if($i != 0){
                $newStr .= ' ';
            }
        }
        return $newStr;
    }
}
$obj = new Solution();
$r = $obj->reverseWords(
    'the sky is blue'
//    "  hello world!  "
//     "a good   example"
);
var_dump($r);
