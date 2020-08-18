<?php
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


$obj = new Trie();
$r = $obj->insert("abc");
var_dump($r);
$r = $obj->search("abc");
var_dump($r);
$r = $obj->startsWith("ab");
var_dump($r);
