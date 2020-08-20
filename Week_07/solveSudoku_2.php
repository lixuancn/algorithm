<?php
//这个思路是把能用的合法数字先算出来，再循环
class Solution {
    private $row = array();
    private $col = array();
    private $block = array();
    private $empty = array();

    /**
     * @param String[][] $board
     * @return NULL
     */
    function solveSudoku(&$board) {
        if(!$board){
            return;
        }
        //初始化：每行、每列、每块可以用的数字
        for($i=0; $i<count($board); $i++){
            for($num=1; $num<10; $num++){
                $this->row[$i][$num] = strval($num);
                $this->col[$i][$num] = strval($num);
                $this->block[$i][$num] = strval($num);
            }
        }
        for($i=0; $i<count($board); $i++){
            for($j=0; $j<count($board[0]); $j++){
                //记录是空位置的坐标，到时候就不需要全棋盘遍历了。
                if($board[$i][$j] == '.'){
                    $this->empty[] = array($i, $j);
                    continue;
                }
                //已有的数字从可用列表中删除
                unset($this->row[$i][$board[$i][$j]]);
                unset($this->col[$j][$board[$i][$j]]);
                $index = intval($i / 3) * 3 + intval($j / 3);
                unset($this->block[$index][$board[$i][$j]]);
            }
        }
        $this->backtrack($board, 0);
        return;
    }

    function backtrack(&$board, $index){
        if($index == count($this->empty)){
            return true;
        }
        list($i, $j) = $this->empty[$index];
        $blockIndex = intval($i / 3) * 3 + intval($j / 3);
        //从可用列表中选出一个数字
        foreach($this->row[$i] as $num){
            if(isset($this->col[$j][$num]) && isset($this->block[$blockIndex][$num])){
                unset($this->row[$i][$num]);
                unset($this->col[$j][$num]);
                unset($this->block[$blockIndex][$num]);
                $board[$i][$j] = $num;
                if($this->backtrack($board, $index+1)){
                    return true;
                }
                //回溯
                $this->row[$i][$num] = $num;
                $this->col[$j][$num] = $num;
                $this->block[$blockIndex][$num] = $num;
            }
        }
        return false;
    }
}
$obj = new Solution();
$a = [
    ["5","3",".",".","7",".",".",".","."],
    ["6",".",".","1","9","5",".",".","."],
    [".","9","8",".",".",".",".","6","."],
    ["8",".",".",".","6",".",".",".","3"],
    ["4",".",".","8",".","3",".",".","1"],
    ["7",".",".",".","2",".",".",".","6"],
    [".","6",".",".",".",".","2","8","."],
    [".",".",".","4","1","9",".",".","5"],
    [".",".",".",".","8",".",".","7","9"]
];
$r = $obj->solveSudoku(
    $a

//    [
//        ["8","3",".",".","7",".",".",".","."],
//        ["6",".",".","1","9","5",".",".","."],
//        [".","9","8",".",".",".",".","6","."],
//        ["8",".",".",".","6",".",".",".","3"],
//        ["4",".",".","8",".","3",".",".","1"],
//        ["7",".",".",".","2",".",".",".","6"],
//        [".","6",".",".",".",".","2","8","."],
//        [".",".",".","4","1","9",".",".","5"],
//        [".",".",".",".","8",".",".","7","9"]
//    ]
);
var_dump($a);
