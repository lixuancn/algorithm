<?php
//DP
class Solution {

    /**
     * @param String $s
     * @return Integer
     */
    function countSubstrings($s) {
        //表示索引i到j这个子字符串是不是回文串
        //dp[i][j] = dp[i+1][j-1] && $s[i]==$s[j]
        $dp = array();
        $dp[0][0] = true;
        for($i=1; $i<strlen($s); $i++){
            $dp[$i][$i] = true;
            $dp[$i-1][$i] = $s[$i-1]==$s[$i] ? true : false;
        }
        for($len=2; $len<strlen($s); $len++){
            for($i=0; $i<strlen($s)-$len; $i++){
                $j = $i + $len;
                $dp[$i][$j] = $dp[$i+1][$j-1] && $s[$i] == $s[$j];
            }
        }
        $sum = 0;
        foreach($dp as $row){
            foreach($row as $col){
                if($col){
                    $sum++;
                }
            }
        }
        return $sum;
    }
}
$ret = (new Solution())->countSubstrings(
    "aaa"
);
var_dump($ret);
