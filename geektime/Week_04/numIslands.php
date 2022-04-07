<?php
//遇到1就计数器+1，然后递归把相邻的1都改成0.
class Solution {

    /**
     * @param String[][] $grid
     * @return Integer
     */
    function numIslands($grid) {
        $y = count($grid);
        if(!$y){
            return 0;
        }
        $x = count($grid[0]);
        if(!$x){
            return 0;
        }
        $count = 0;
        for($i=0; $i<$y; $i++){
            for($j=0; $j<$x; $j++){
                if($grid[$i][$j] == '1'){
                    $count++;
                    $grid = $this->dfs($grid, $i, $j, $y, $x);
                }
            }
        }
        return $count;
    }

    private function dfs($grid, $i, $j, $y, $x){
        if($i < 0 || $j < 0 || $i >= $y || $j >= $x || $grid[$i][$j] != '1'){
            return $grid;
        }
        $grid[$i][$j] = '0';
        $grid = $this->dfs($grid, $i+1, $j, $y, $x);
        $grid = $this->dfs($grid, $i-1, $j, $y, $x);
        $grid = $this->dfs($grid, $i, $j+1, $y, $x);
        $grid = $this->dfs($grid, $i, $j-1, $y, $x);
        return $grid;
    }
}

$ret = (new Solution())->numIslands([
    ['1','1','0','0','0'],
    ['1','1','0','0','0'],
    ['0','0','1','0','0'],
    ['0','0','0','1','1']
]);
var_dump($ret);
