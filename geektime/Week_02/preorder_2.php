<?php
//N叉树 - 前序遍历 - 方法一：递归
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
    function preorder($root) {
        $list = array();
        $this->helper($root, $list);
        return $list;
    }

    function helper($node, &$list){
        $list[] = $node->val;
        foreach($node->children as $n){
            if($n){
                $this->helper($n, $list);
            }
        }
    }
}
