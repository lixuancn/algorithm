<?php
//DP
class Solution {

    /**
     * @param Integer[][] $triangle
     * @return Integer
     */
    function minimumTotal($triangle) {
        if(!$triangle){
            return 0;
        }
        for($i=count($triangle)-2; $i>=0; $i--){
            for($j=0; $j<=$i; $j++){
                $triangle[$i][$j] = min($triangle[$i+1][$j], $triangle[$i+1][$j+1]) + $triangle[$i][$j];
            }
        }
        return $triangle[0][0];
    }
}
$ret = (new Solution())->minimumTotal(
    [
        [2],
        [3,4],
        [6,5,7],
        [4,1,8,3]
    ]
);
var_dump($ret);
