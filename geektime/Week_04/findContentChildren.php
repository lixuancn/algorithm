<?php
//贪心
class Solution {

    /**
     * @param Integer[] $g
     * @param Integer[] $s
     * @return Integer
     */
    function findContentChildren($g, $s) {
        sort($g);
        sort($s);
        $ret = 0;
        $i = $j = 0;
        while ($i < count($g) && $j < count($s)){
            if($g[$i] <= $s[$j]){
                $ret++;
                $i++;
                $j++;
            }else{
                $j++;
            }
        }
        return $ret;
    }
}

$ret = (new Solution())->findContentChildren(
    [1,2,3], [1,1]
);
var_dump($ret);
