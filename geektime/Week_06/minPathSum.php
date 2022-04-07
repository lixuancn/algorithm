<?php
//DP
class Solution {

    /**
     * @param Integer[][] $grid
     * @return Integer
     */
    function minPathSum($grid) {
        if(!$grid || count($grid[0]) < 0) return 0;
        if(count($grid[0]) < 1){
            return $grid[0][0];
        }
        $dp = $grid;
        $max = pow(2, 31);
        for($i=count($grid)-1; $i>=0; $i--){
            for($j=count($grid[$i])-1; $j>=0; $j--){
                //右下角
                if(!isset($dp[$i+1][$j]) && !isset($dp[$i][$j+1])){
                    continue;
                }
                //下边没有只看右边
                else if(!isset($dp[$i+1][$j])){
                    $dp[$i][$j] = $dp[$i][$j] + $dp[$i][$j+1];
                }
                //右边没有只看下边
                else if(!isset($dp[$i][$j+1])){
                    $dp[$i][$j] = $dp[$i][$j] + $dp[$i+1][$j];
                }else{
                    $dp[$i][$j] = min($dp[$i+1][$j], $dp[$i][$j+1]) + $grid[$i][$j];
                }
            }
        }
        return $dp[0][0];
    }
}
$ret = (new Solution())->minPathSum(
    [
        [1,3,1],
        [1,5,1],
        [4,2,1]
    ]
);
var_dump($ret);
