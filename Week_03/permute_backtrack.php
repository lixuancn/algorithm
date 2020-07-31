<?php
//回溯
class Solution {

    private $ret = array();

    /**
     * @param Integer[] $nums
     * @return Integer[][]
     */
    function permute($nums) {
        $ans = array();
        foreach($nums as $k => $num){
            $ans[] = $num;
        }
        $this->backtrack(count($nums), $ans, 0);
        return $this->ret;
    }

    function backtrack($n, $ans, $first){
        if($n == $first){
            $this->ret[] = $ans;
        }
        for($i=$first; $i<$n; $i++){
            $tmp = $ans[$i];
            $ans[$i] = $ans[$first];
            $ans[$first] = $tmp;
            $this->backtrack($n, $ans, $first+1);
            $tmp = $ans[$i];
            $ans[$i] = $ans[$first];
            $ans[$first] = $tmp;
        }
    }
}

$ret = (new Solution())->permute([1,2,3]);
var_dump($ret);
