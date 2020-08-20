<?php
//递归，剪枝
class Solution {
    private $ret = array();
    /**
     * @param Integer $n
     * @return String[]
     */
    function generateParenthesis($n) {
        $this->generate(0, 0, $n, '');
        return $this->ret;
    }

    private function generate($left, $right, $n, $s){
        if($left == $n && $right == $n){
            $this->ret[] = $s;
            return;
        }
        if($left < $n){
            $this->generate($left+1, $right, $n, $s.'(');
        }
        if($right < $left){
            $this->generate($left, $right+1, $n, $s.')');
        }
    }
}
$obj = new Solution();
$r = $obj->generateParenthesis(
    3
);
var_dump($r);
