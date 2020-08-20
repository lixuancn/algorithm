<?php
//递归
class Solution {
    /**
     * @param String[][] $board
     * @return NULL
     */
    function solveSudoku(&$board) {
        if(!$board){
            return;
        }
        $this->dfs($board);
        return;
    }

    function dfs(&$board){
        for($i=0; $i<count($board); $i++){
            for($j=0; $j<count($board); $j++) {
                if($board[$i][$j] != '.') {
                    continue;
                }
                for($num=1; $num<=9; $num++){
                    $num = strval($num);
                    if($this->isValidSudoku($board, $i, $j, $num)){
                        $board[$i][$j] = $num;
                        if($this->dfs($board)){
                            return true;
                        }
                        $board[$i][$j] = '.';
                    }
                }
                return false;
            }
        }
        return true;
    }

    /**
     * @param String[][] $board
     * @return Boolean
     */
    function isValidSudoku($board, $i, $j, $num) {
        for($c=0; $c<9; $c++) {
            if ($board[$c][$j] != '.' && $board[$c][$j] == $num) {
                return false;
            }
            if ($board[$i][$c] != '.' && $board[$i][$c] == $num) {
                return false;
            }
            $k1 = intval($i / 3) * 3 + intval($c / 3);
            $k2 = intval($j / 3) * 3 + intval($c % 3);
            if ($board[$k1][$k2] != '.' && $board[$k1][$k2] == $num) {
                return false;
            }
        }
        return true;
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
