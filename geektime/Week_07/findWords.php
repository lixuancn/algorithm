<?php
//前缀树
class Trie {
    private $isEnd = false;
    private $next = array();
    /**
     * Initialize your data structure here.
     */
    function __construct() {

    }

    /**
     * Inserts a word into the trie.
     * @param String $word
     * @return NULL
     */
    function insert($word) {
        $cur = $this;
        for($i=0; $i<strlen($word); $i++){
            $w = $word[$i];
            if(!isset($cur->next[$w])){
                $cur->next[$w] = new Trie();
            }
            $cur = $cur->next[$w];
        }
        $cur->isEnd = true;
        return;
    }

    /**
     * Returns if the word is in the trie.
     * @param String $word
     * @return Boolean
     */
    function search($word) {
        $node = $this->searchPrefix($word);
        return $node && $node->isEnd === true ? true : false;
    }

    /**
     * Returns if there is any word in the trie that starts with the given prefix.
     * @param String $prefix
     * @return Boolean
     */
    function startsWith($prefix) {
        $node = $this->searchPrefix($prefix);
        return $node ? true : false;
    }

    private function searchPrefix($word) {
        $cur = $this;
        for($i=0; $i<strlen($word); $i++){
            if(!isset($cur->next[$word[$i]])){
                return false;
            }
            $cur = $cur->next[$word[$i]];
        }
        return $cur;
    }
}

class Solution {
    private $trie = array();
    private $ret = array();
    /**
     * @param String[][] $board
     * @param String[] $words
     * @return String[]
     */
    function findWords($board, $words) {
        if(!$words || !$board || !count($board[0])){
            return [];
        }
        $this->trie = new Trie();
        //构造树
        foreach($words as $word){
            $this->trie->insert($word);
        }
        for($i=0; $i<count($board); $i++){
            for($j=0; $j<count($board[$i]); $j++) {
                if($this->trie->startsWith($board[$i][$j])){
                    $this->dfs($board, $i, $j, count($board), count($board[$i]), "");
                }
            }
        }
        return array_keys($this->ret);
    }

    function dfs($board, $i, $j, $m, $n, $w){
        $w = $w.$board[$i][$j];
        if(!$this->trie->startsWith($w)){
            return;
        }
        if($this->trie->search($w)){
            $this->ret[$w] = 1;
        }
        $tmp = $board[$i][$j];
        $board[$i][$j] = '@';
        $iiList = array(-1, 1, 0, 0);
        $jjList = array(0, 0, -1, 1);
        for($x=0; $x<4; $x++){
            $newI = $i + $iiList[$x];
            $newJ = $j + $jjList[$x];
            if($newI < 0 || $newI >= $m || $newJ < 0 || $newJ >= $n){
                continue;
            }
            if($board[$newI][$newJ] == '@'){
                continue;
            }
            $this->dfs($board, $newI, $newJ, $m, $n, $w);
        }
        $board[$i][$j] = $tmp;
    }
}

$obj = new Solution();
$r = $obj->findWords(
//    [
//    ['o','a','a','n'],
//    ['e','t','a','e'],
//    ['i','h','k','r'],
//    ['i','f','l','v']
//], ["oath","pea","eat","rain"]
[["a", "a"]], ["a"]
);
var_dump($r);
