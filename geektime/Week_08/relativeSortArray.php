<?php
class Solution {

    /**
     * @param Integer[] $arr1
     * @param Integer[] $arr2
     * @return Integer[]
     */
    function relativeSortArray($arr1, $arr2) {
        $count = array_fill(0, 1000, 0);
        foreach($arr1 as $v){
            $count[$v]++;
        }
        $i = 0;
        foreach($arr2 as $v){
            while($count[$v]){
                $arr1[$i++] = $v;
                $count[$v]--;
            }
        }
        foreach($count as $v=>$c){
            while($c){
                $arr1[$i++] = $v;
                $c--;
            }
        }
        return $arr1;
    }
}

$obj = new Solution();
$r = $obj->relativeSortArray([2,3,1,3,2,4,6,19,9,2,7], [2,1,4,3,9,6]);
var_dump($r);
