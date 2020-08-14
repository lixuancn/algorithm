<?php
//dp
//每个点可能偷，也可能不偷，偷就用$a[$i][0]表示，不偷$a[$i][1]表示（这时才可以+$nums[$i]）
class Solution {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function rob($nums) {
        if(!$nums){
            return 0;
        }
        $dp = array();
        $dp[0][0] = $nums[0];
        $dp[0][1] = 0;
        for($i=1; $i<count($nums); $i++){
            $dp[$i][0] = $dp[$i-1][1] + $nums[$i];
            $dp[$i][1] = max($dp[$i-1][0], $dp[$i-1][1]);
        }
        return max($dp[count($nums)-1][0], $dp[count($nums)-1][1]);
    }
}
$ret = (new Solution())->rob(
//    [1,2,3,1]
    [2,7,9,3,1]
);
var_dump($ret);
