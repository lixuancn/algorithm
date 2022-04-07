<?php
//DP
class Solution {

    /**
     * @param String[][] $matrix
     * @return Integer
     */
    function maximalSquare($matrix) {
        if(!$matrix || count($matrix) < 1 || !count($matrix[0])){
            return 0;
        }
        $rows = count($matrix);
        $cols = count($matrix[0]);
        $dp = array();
        $maxSide = 0;
        for($i=0; $i<$rows; $i++){
            for($j=0; $j<$cols; $j++){
                $dp[$i][$j] = 0;
                if($matrix[$i][$j] == 1){
                    if($i==0 || $j==0){
                        $dp[$i][$j] = 1;
                    }else{
                        $dp[$i][$j] = 1 + min($dp[$i-1][$j-1], $dp[$i-1][$j], $dp[$i][$j-1]);
                    }
                    $maxSide = max($maxSide, $dp[$i][$j]);
                }
            }
        }
        return $maxSide * $maxSide;
    }
}
$ret = (new Solution())->maximalSquare(
    [
        [1,0,1,0,0],
        [1,0,1,1,1],
        [1,1,1,1,1],
        [1,0,0,1,0],
    ]
);
var_dump($ret);
