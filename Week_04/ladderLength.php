<?php
//广度优先
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
        $pregWordList = array();
        foreach($wordList as $w){
            if($w == $beginWord){
                continue;
            }
            for($i=0; $i<strlen($w); $i++) {
                $preg = $w;
                $preg[$i] = '*';
                $pregWordList[$preg][] = $w;
            }
        }
        $visitedList = array();
        $queue = array();
        $queue[] = array($beginWord, 1);
        while($queue){
            list($word, $step) = array_shift($queue);
            var_dump($word);
            for($i=0; $i<strlen($word); $i++){
                $preg = $word;
                $preg[$i] = '*';
                if(!isset($pregWordList[$preg])){
                    continue;
                }
                foreach($pregWordList[$preg] as $k => $preg){
                    $nextWord = $preg;
                    if($nextWord == $endWord){
                        return $step + 1;
                    }
                    if(isset($visitedList[$nextWord])){
                        continue;
                    }
                    $visitedList[$nextWord] = 1;
                    $queue[] = array($nextWord, $step + 1);
                    unset($pregWordList[$preg][$k]);
                }
            }
        }
        return 0;
    }
}

$ret = (new Solution())->ladderLength('hit', 'cog', ["hot","dot","tog","cog"]);
var_dump($ret);
