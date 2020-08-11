<?php
//递归+记忆
class Solution {

    private $triangle;
    private $cache = array();
    /**
     * @param Integer[][] $triangle
     * @return Integer
     */
    function minimumTotal($triangle) {
        if(!$triangle){
            return 0;
        }
        $this->triangle = $triangle;
        return $this->helper(0, 0);
    }

    private function helper($i, $j)
    {
        if(isset($this->cache[$i][$j])){
            return $this->cache[$i][$j];
        }
        if($i == count($this->triangle)-1){
            return $this->triangle[$i][$j];
        }
        $left = $this->helper($i+1, $j);
        $right = $this->helper($i+1, $j+1);
        $result = min($left, $right) + $this->triangle[$i][$j];
        $this->cache[$i][$j] = $result;
        return $result;
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
