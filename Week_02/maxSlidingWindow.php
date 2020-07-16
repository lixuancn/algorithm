<?php
//堆的方式
/**
 * Definition for a Node.
 * class Node {
 *     public $val = null;
 *     public $children = null;
 *     function __construct($val = 0) {
 *         $this->val = $val;
 *         $this->children = array();
 *     }
 * }
 */

class Solution {
    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return Integer[]
     */
    function maxSlidingWindow($nums, $k) {
        if(!$k || !$nums){
            return [];
        }
        for($i=0; $i<count($nums)-$k+1; $i++){
            $queue = new \SplPriorityQueue();
            for($j=$i; $j<$i+$k; $j++){
                $queue->insert($nums[$j], $nums[$j]);
            }
            $ret[] = $queue->top();
        }
        return $ret;
    }
}

$r = (new Solution())->maxSlidingWindow([1,3,-1,-3,5,3,6,7], 3);
var_dump($r);
