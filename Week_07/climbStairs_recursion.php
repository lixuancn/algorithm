<?php
//递归，剪纸
class Solution {
    private $cache = array();
    /**
     * @param Integer $n
     * @return Integer
     */
    function climbStairs($n) {
        if($n == 0) return 0;
        if($n == 1) return 1;
        if($n == 2) return 2;
        if(isset($this->cache[$n])){
           return $this->cache[$n];
        }
        $r = $this->climbStairs($n-1) + $this->climbStairs($n-2);
        $this->cache[$n] = $r;
        return $r;
    }
}
$obj = new Solution();
$r = $obj->climbStairs(
    3
);
var_dump($r);
