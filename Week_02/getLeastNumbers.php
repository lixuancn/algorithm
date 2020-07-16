<?php
//最小的K个数，用最大堆的方式，堆元素总数为k个。大的都出堆，那最终堆里的就是最小的K个。
//插入堆是O(lonK)，所以最终是O(NlogK)
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
     * @param Integer[] $arr
     * @param Integer $k
     * @return Integer[]
     */
    function getLeastNumbers($arr, $k) {
        if(!$k){
            return [];
        }
        $queue = new \SplPriorityQueue();
        for($i=0; $i<$k; $i++){
            $queue->insert($arr[$i], $arr[$i]);
        }
        for($i=$k; $i<count($arr); $i++){
            if($queue->top() > $arr[$i]){
                $queue->extract();
                $queue->insert($arr[$i], $arr[$i]);
            }
        }
        $ret = array();
        while($queue->valid()){
            $ret[] = $queue->current();
            $queue->next();
        }
        return $ret;
    }
}

$r = (new Solution())->getLeastNumbers([0,0,0,2,0,5], 0);
var_dump($r);
