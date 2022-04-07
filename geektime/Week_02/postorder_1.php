<?php
//N叉树 - 后序遍历 - 方法一：递归

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
     * @return integer[]
     */
    function postorder($root) {
        $ret = array();
        $this->handler($root, $ret);
        return $ret;
    }
    
    private function handler($root, &$ret){
        foreach($root->children as $node){
            if($node){
                $this->handler($node, $ret);
            }
        }
        $ret[] = $root->val;
    }
}
