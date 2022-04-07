<?php
//BFS
class Solution {

    /**
     * @param Integer[][] $grid
     * @return Integer
     */
    function shortestPathBinaryMatrix($grid)
    {
        if (!$grid || $grid[0][0] == 1 || $grid[count($grid) - 1][count($grid[0]) - 1] == 1) {
            return -1;
        }
        $n = count($grid);
        if ($n <= 2) {
            return $n;
        }
        $queue = array();
        $queue[] = array(0, 0, 2);
        while ($queue) {
            $node = array_shift($queue);
            list($i, $j, $step) = $node;
            foreach ([[-1, -1], [0, -1], [-1, 0], [1, 0], [0, 1], [1, 1], [-1, 1], [1, -1]] as $pos) {
                $newI = $i + $pos[0];
                $newJ = $j + $pos[1];
                if ($newI >= 0 && $newI < $n && $newJ >= 0 && $newJ < $n && $grid[$newI][$newJ] != 1) {
                    if ($newI == $n - 1 && $newJ == $n - 1) {
                        return $step;
                    }
                    $queue[] = array($newI, $newJ, $step + 1);
                    //实现visited的功能
                    $grid[$newI][$newJ] = 1;
                }
            }
        }
        return -1;
    }
}

$obj = new Solution();
$r = $obj->shortestPathBinaryMatrix(
//    [[0,1],[1,0]]
//    [[0,1],[1,1],[1,0]]
    [[0,0,0],[1,1,0],[1,1,0]]
);
var_dump($r);
