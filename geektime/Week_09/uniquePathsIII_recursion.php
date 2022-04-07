<?php
//递归
class Solution {
    private $ans = 0;
    private $si;
    private $sj;
    private $di;
    private $dj;
    private $grid;

    /**
     * @param Integer[][] $grid
     * @return Integer
     */
    function uniquePathsIII($grid) {
        if(!$grid || !count($grid[0])) return 0;
        $m = count($grid);
        $n = count($grid[0]);
        //找到起始点 和 可以走的格子数量
        $enableCell = 0;
        for($i=0; $i<$m; $i++) {
            for ($j=0; $j<$n; $j++) {
                if($grid[$i][$j] == -1){
                    continue;
                }
                $enableCell++;
                if($grid[$i][$j] == 1){
                    $this->si = $i;
                    $this->sj = $j;
                }else if($grid[$i][$j] == 2){
                    $this->di = $i;
                    $this->dj = $j;
                }
            }
        }
        $this->grid = $grid;
        $this->dfs($m, $n, $this->si, $this->sj, $enableCell);
        return $this->ans;
    }

    private function dfs($m, $n, $i, $j, $enableCell){
        $enableCell--;
        if($enableCell < 0) return;
        if($this->di == $i && $this->dj == $j){
            if($enableCell == 0){
                $this->ans++;
            }
            return;
        }
        $this->grid[$i][$j] = 3;
        foreach(array(array(0,-1), array(-1,0), array(1,0), array(0,1)) as $pos){
            $newI = $i + $pos[0];
            $newJ = $j + $pos[1];
            if($newI>=0 && $newI<$m && $newJ>=0 && $newJ<$n && $this->grid[$newI][$newJ]%2==0){
                $this->dfs($m, $n, $newI, $newJ, $enableCell);
            }
        }
        $this->grid[$i][$j] = 0;
    }
}
$obj = new Solution();
$r = $obj->uniquePathsIII(
    [[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
);
var_dump($r);
