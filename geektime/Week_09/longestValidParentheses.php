<?php
//DP
class Solution {
    /**
     * @param String $s
     * @return Integer
     */
    function longestValidParentheses($s) {
        $n = strlen($s);
        $max = 0;
        //dp[i]表示以i结尾的子串长度，所以只要i==(，dp[i]就为0
        $dp = array();
        $dp[0] = 0;
        for($i=1; $i<$n; $i++){
            if($s[$i] == '('){
                $dp[$i] = 0;
            }else if($s[$i] == ')'){
                if($s[$i-1] == '('){
                    $dp[$i] = $dp[$i-2] + 2;
                    //>=0的判断是要有的，不然负数就从尾巴开始数了
                }elseif($s[$i-1] == ')' && $i-1-$dp[$i-1]>= 0 && $s[$i-1-$dp[$i-1]]=='('){
                    $dp[$i] = $dp[$i-1] + 2 + $dp[$i-1-$dp[$i-1]-1];
                }
            }
            $max = max($max, $dp[$i]);
        }
        return $max;
    }
}
$obj = new Solution();
$r = $obj->longestValidParentheses(
//    '(()'
//    ')()())'
//    ')('
//    "()(()"
    "(()))())(1"
);
var_dump($r);
