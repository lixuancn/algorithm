<?php
//dp
//上一个基础上优化：二维合成一维.
//dp变为：dp[i] = 从0到i 且 i必偷的最大值。
//dp[i] = max(dp[i-1], dp[i-2]+nums[i], dp[i-3]+nums[i], dp[i-4]+nums[i]...)
//可证明dp[i] 就是= max(dp[i-1], dp[i-2]+nums[i])
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
        $dp[0] = $nums[0];
        $dp[1] = max($nums[0], $nums[1]);
        $ans = max($dp[0], $dp[1]);
        for($i=2; $i<count($nums); $i++){
            $dp[$i] = max($dp[$i-1], $dp[$i-2] + $nums[$i]);
            $ans = max($dp[$i], $ans);
        }
        return $ans;
    }
}
