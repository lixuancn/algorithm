<?php
//BFS
class Solution
{
    /**
     * @param Integer[][] $board
     * @return Integer
     */
    function slidingPuzzle($board)
    {
        if (!$board) {
            return -1;
        }
        $m = 2;
        $n = 3;
        $queue = new SplPriorityQueue();
        $pos = $this->getZero($board);
        $pos[] = $board;
        $pos[] = 0;
        $queue->insert($pos, 0);
        $bString = $this->getString($board);
        $visited = array($bString=>1);
        while ($queue->count()) {
            $newQueue = new SplPriorityQueue();
            while ($queue->count() && $node = $queue->extract()){
                list($i, $j, $b, $step) = $node;
                if($i == $m-1 && $j == $n-1){
                    if($this->getString($b) == '123450'){
                        return $step;
                    }
                }
                foreach ([[0, 1], [0, -1], [-1, 0], [1, 0]] as $pos) {
                    $newI = $i + $pos[0];
                    $newJ = $j + $pos[1];
                    if ($newI >= 0 && $newI < $m && $newJ >= 0 && $newJ < $n) {
                        $bTmp = $b;
                        $bTmp[$i][$j] = $bTmp[$newI][$newJ];
                        $bTmp[$newI][$newJ] = 0;
                        $bString = $this->getString($bTmp);
                        if(!isset($visited[$bString])){
                            $visited[$bString] = 1;
                            $next = array($newI, $newJ, $bTmp, $step+1);
                            $newQueue->insert($next, 0);
                        }
                    }
                }
            }
            $queue = $newQueue;
        }
        return -1;
    }

    function getZero($board)
    {
        for ($i = 0; $i < 2; $i++) {
            for ($j = 0; $j < 3; $j++) {
                if ($board[$i][$j] == 0) {
                    return array($i, $j);
                }
            }
        }
    }

    function getString($board){
        $s = '';
        for ($i = 0; $i < 2; $i++) {
            for ($j = 0; $j < 3; $j++) {
                $s .= $board[$i][$j];
            }
        }
        return $s;
    }
}

$obj = new Solution();
$r = $obj->slidingPuzzle(
//    [[1,2,3],[4,0,5]]
    [[4,1,2],[5,0,3]]
);
var_dump($r);
