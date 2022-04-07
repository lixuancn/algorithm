<?php

/**
 * 递归，O(2^n)，使用记忆递归，优化到O(n)
 * Class Solution
 */
class Solution {

    private $memo = array();

    /**
     * @param Integer $n
     * @return Integer
     */
    function climbStairs($n) {
        if(isset($this->memo[$n])){
            return $this->memo[$n];
        }
        if($n == 1) return 1;
        if($n == 2) return 2;
        $ret = $this->climbStairs($n-1) + $this->climbStairs($n-2);
        $this->memo[$n] = $ret;
        return $ret;
    }
}
$obj = new Solution();
$nums = $obj->climbStairs(3);
print_r($nums);
