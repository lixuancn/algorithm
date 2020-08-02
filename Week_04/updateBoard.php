<?php

class Solution {

    /**
     * @param String[][] $board
     * @param Integer[] $click
     * @return String[][]
     */
    function updateBoard($board, $click)
    {
        list($i, $j) = $click;
        $row = count($board);
        $col = count($board[0]);
        if ($board[$i][$j] == "M") {
            $board[$i][$j] = "X";
            return $board;
        }
        $board = $this->dfs($i, $j, $board, $row, $col);
        return $board;
    }

    // 计算空白快周围的***
    function cal($i, $j, $row, $col, $board){
        $res = 0;
        foreach([1, -1, 0] as $x){
            foreach([1, -1, 0] as $y){
                if($x == 0 && $y == 0){
                    continue;
                }
                if($i+$x>=0 && $i+$x<$row && $j+$y>=0 && $j+$y<$col && $board[$i+$x][$j+$y] == 'M'){
                    $res++;
                }
            }
        }
        return $res;
    }

    function dfs($i, $j, $board, $row, $col){
        $num = $this->cal($i, $j, $row, $col, $board);
        if($num > 0){
            $board[$i][$j] = strval($num);
            return $board;
        }
        $board[$i][$j] = 'B';
        foreach([1,-1,0] as $x) {
            foreach([1,-1,0] as $y) {
                if($x == 0 && $y == 0) {
                    continue;
                }
                $nextI = $i + $x;
                $nextJ = $j + $y;
                if($nextI >= 0 && $nextI < $row && $nextJ >= 0 && $nextJ < $col && $board[$nextI][$nextJ] == 'E') {
                    $board = $this->dfs($nextI, $nextJ, $board, $row, $col);
                }
            }
        }
        return $board;
    }
}

$ret = (new Solution())->updateBoard(
    [['E', 'E', 'E', 'E', 'E'],
        ['E', 'E', 'M', 'E', 'E'],
        ['E', 'E', 'E', 'E', 'E'],
        ['E', 'E', 'E', 'E', 'E']], [3,0]
);
var_dump($ret);
