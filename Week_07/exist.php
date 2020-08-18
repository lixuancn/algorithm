<?php
//递归

class Solution {
    private $trie = array();
    private $ret = false;
    /**
     * @param String[][] $board
     * @param String $word
     * @return Boolean
     */
    function exist($board, $word) {
        if(!$board || !$word || !count($board[0])){
            return false;
        }
        for($i=0; $i<count($board); $i++){
            for($j=0; $j<count($board[0]); $j++) {
                $this->dfs($board, $word, count($board), count($board[0]), $i, $j, "");
                if($this->ret){
                    return $this->ret;
                }
            }
        }
        return false;
    }

    function dfs($board, $word, $m, $n, $i, $j, $w){
        $w .= $board[$i][$j];
        if($w == $word){
            $this->ret = true;
            return;
        }
        for($s=0; $s<strlen($w); $s++){
            if($w[$s] != $word[$s]){
                return;
            }
        }
        $tmp = $board[$i][$j];
        $board[$i][$j] = '@';
        $iiList = array(-1, 1, 0, 0);
        $jjList = array(0, 0, -1, 1);
        for($x=0; $x<4; $x++) {
            $newI = $i + $iiList[$x];
            $newJ = $j + $jjList[$x];
            if ($newI < 0 || $newI >= $m || $newJ < 0 || $newJ >= $n) {
                continue;
            }
            if($board[$newI][$newJ] == '@'){
                continue;
            }
            $this->dfs($board, $word, $m, $n, $newI, $newJ, $w);
        }
        $board[$i][$j] = $tmp;
    }
}

$obj = new Solution();
$r = $obj->exist(
    [["A"]]
    ,"A"
);
var_dump($r);
