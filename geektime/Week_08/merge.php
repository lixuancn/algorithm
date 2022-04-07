<?php

class Solution {

    /**
     * @param Integer[][] $intervals
     * @return Integer[][]
     */
    function merge($intervals) {
        usort($intervals, function ($a, $b){
            return $a[0] > $b[0];
        });
        $ret = array();
        foreach($intervals as $interval){
            $count = count($ret);
            //当前区间和上个区间不重合
            if(!$ret || $interval[0] > $ret[$count-1][1]){
                $ret[] = $interval;
            }else{
                $ret[$count-1][1] = max($interval[1], $ret[$count-1][1]);
            }
        }
        return $ret;
    }
}
