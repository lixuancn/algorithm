<?php
//翻转一棵树 - 循环

/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($value) { $this->val = $value; }
 * }
 */
class Solution {

    /**
     * @param TreeNode $root
     * @return TreeNode
     */
    function invertTree($root) {
        if(!$root){
            return $root;
        }
        $queue = array();
        array_unshift($queue, $root);
        while(count($queue)){
            $node = array_pop($queue);
            $tmp = $node->right;
            $node->right = $node->left;
            $node->left = $tmp;
            if($node->left){
                array_unshift($queue, $node->left);
            }
            if($node->right){
                array_unshift($queue, $node->right);
            }
        }
        return $root;
    }
}
