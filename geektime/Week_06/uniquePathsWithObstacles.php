<?php
//dp
class Solution {
    /**
     * @param Integer[][] $obstacleGrid
     * @return Integer
     */
    function uniquePathsWithObstacles($obstacleGrid) {
        if(!count($obstacleGrid) || !count($obstacleGrid[0])){
            return 0;
        }
        $ret = array();
        for($i=0; $i<count($obstacleGrid[0]) && $obstacleGrid[0][$i] != 1; $i++){
            $ret[0][$i] = 1;
        }
        for($i=0; $i<count($obstacleGrid) && $obstacleGrid[$i][0] != 1; $i++){
            $ret[$i][0] = 1;
        }
        var_dump($ret);
        for($i=1; $i<count($obstacleGrid); $i++){
            for($j=1; $j<count($obstacleGrid[0]); $j++){
                if($obstacleGrid[$i][$j] == 1){
                    $ret[$i][$j] = 0;
                    continue;
                }
                $ret[$i][$j] = $ret[$i-1][$j] + $ret[$i][$j-1];
            }
        }
        return intval($ret[count($obstacleGrid)-1][count($obstacleGrid[0])-1]);
    }
}
$ret = (new Solution())->uniquePathsWithObstacles(
    [
        [0,1,0],
        [0,1,0],
        [0,0,0]
    ]
);
var_dump($ret);

