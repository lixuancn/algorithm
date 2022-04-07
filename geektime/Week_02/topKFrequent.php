<?php
//前K个高频元素。堆的方式

class Solution {
    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return Integer[]
     */
    function topKFrequent($nums, $k) {
        if(!$nums || !$k){
            return [];
        }
        $count = array();
        foreach($nums as $num){
            $count[$num] = !isset($count[$num]) ? 1 : ++$count[$num];
        }
        $heap = new \Heap();
        foreach($count as $n => $c){
            $heap->insert($n, $c);
            if($heap->count() > $k){
                $heap->extract();
            }
        }
        $ret = array();
        while($heap && $heap->valid() && $heap->count() > 0){
            $ret[] = $heap->current();
            $heap->next();
        }
        return $ret;
    }
}

class Heap extends \SplPriorityQueue{
    public function compare($priority1, $priority2)
    {
        return $priority2 - $priority1;
    }
}

$r = (new Solution())->topKFrequent([1,1,1,2,2,3], 2);
var_dump($r);
