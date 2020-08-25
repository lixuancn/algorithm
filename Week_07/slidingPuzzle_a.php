<?php
//A*
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
                            $score = $this->getScore($bTmp);
                            $newQueue->insert($next, $score);
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

    function getScore($board)
    {
        $result = 0;
        for ($i = 0; $i < 2; $i++) {
            for ($j = 0; $j < 3; $j++) {
                if ($board[$i][$j] == 0) continue;
                //二维坐标曼哈顿距离计算：res = |x - i| + |y - j|
                $val = $board[$i][$j] - 1;
                $x = intval($val / 3);
                $y = intval($val % 3);
                $result += abs($x - $i) + abs($y - $j);
            }
        }
        return $result;
    }
}

$obj = new Solution();
$r = $obj->slidingPuzzle(
//    [[1,2,3],[4,0,5]]
    [[4,1,2],[5,0,3]]
);
var_dump($r);
