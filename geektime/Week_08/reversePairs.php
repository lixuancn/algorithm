<?php
//归并
class Solution {
    private $nums = array();
    /**
     * @param Integer[] $nums
     * @return Integer
     */
    function reversePairs($nums) {
        if(!$nums) return 0;
        $this->nums = $nums;
        return $this->mergeSort(0, count($nums)-1);
    }

    private function mergeSort($start, $end)
    {
        if ($start >= $end) return 0;
        $mid = $start + intval(($end - $start) / 2);
        $res = 0;
        $res += $this->mergeSort($start, $mid);
        $res += $this->mergeSort($mid + 1, $end);
        $res += $this->merge($start, $mid, $end);
        return $res;
    }

    private function merge($start, $mid, $end){
        $count = 0;
        $a = array_fill(0, $end - $start + 1, 0);
        $i = $start;
        $j = $mid + 1;
        while ($i <= $mid && $j <= $end) {
            if ($this->nums[$i] > 2 * $this->nums[$j]) {
                $count += $mid - $i + 1;
                $j++;
            } else {
                $i++;
            }
        }
        $i = $start;
        $j = $mid + 1;
        $index = 0;
        while ($i <= $mid && $j <= $end) {
            if ($this->nums[$i] > $this->nums[$j]) {
                $a[$index++]  = $this->nums[$j++];
            } else {
                $a[$index++]  = $this->nums[$i++];
            }
        }
        while ($i <= $mid) {
            $a[$index++]  = $this->nums[$i++];
        }
        while ($j <= $end) {
            $a[$index++]  = $this->nums[$j++];
        }
        foreach($a as $v){
            $this->nums[$start] = $v;
            $start++;
        }
        return $count;
    }
}

$obj = new Solution();
$r = $obj->reversePairs (
    [1,3,2,3,1]
//    [2,4,3,5,1]
);
var_dump($r);
