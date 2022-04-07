<?php
//双向BFS
class Solution {

    /**
     * @param String $beginWord
     * @param String $endWord
     * @param String[] $wordList
     * @return Integer
     */
    function ladderLength($beginWord, $endWord, $wordList) {
        if(!$beginWord || !$endWord || !$wordList || !in_array($endWord, $wordList)){
            return 0;
        }
        if($beginWord == $endWord){
            return 0;
        }
        $visited = array();
        $beginSet = array($beginWord=>1);
        $endSet = array($endWord=>1);
        $wordSet = array();
        foreach($wordList as $word){
            $wordSet[$word] = 1;
        }
        $step = 1;
        while($beginSet && $endSet){
            if(count($beginSet) > count($endSet)){
                $tmp = $beginSet;
                $beginSet = $endSet;
                $endSet = $tmp;
            }
            $tmpSet = array();
            foreach($beginSet as $word => $v){
                for($i=0; $i<strlen($word); $i++){
                    for($j=0; $j<26; $j++){
                        $w = $word;
                        $c = chr(ord('a') + $j);
                        $w[$i] = $c;
                        if(isset($endSet[$w])){
                            return $step+1;
                        }
                        if(!isset($visited[$w]) && isset($wordSet[$w])){
                            $tmpSet[$w] = 1;
                            $visited[$w] = 1;
                        }
                    }
                }
            }
            $step++;
            $beginSet = $tmpSet;
        }
        return 0;
    }
}

$obj = new Solution();
$r = $obj->ladderLength('hot', 'cot', ["cot"]);
var_dump($r);
