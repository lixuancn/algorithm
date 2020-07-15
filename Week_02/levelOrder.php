<?php
//N叉树 - 层序遍历 - 广度优先
//使用队列，而栈是深度优先。
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
     * @param Node $root
     * @return integer[][]
     */
    function levelOrder($root) {
        if(!$root){
            return [];
        }
        $queue = array();
        $ret = array();
        array_push($queue, $root);
        while(count($queue) > 0){
            $list = array();
            $count = count($queue);
            while($count--){
                $node = array_shift($queue);
                $list[] = $node->val;
                foreach($node->children as $children){
                    if($children){
                        array_push($queue, $children);
                    }
                }
            }
            $ret[] = $list;
        }
        return $ret;
    }
}
