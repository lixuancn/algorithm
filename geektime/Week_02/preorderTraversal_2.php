<?php
//前序遍历 - 方法二：栈

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
     * @return Integer[]
     */
    function preorderTraversal($root) {
        $stack = array();
        $list = array();
        array_push($stack, $root);
        while(count($stack)){
            $node = array_pop($stack);
            $list[] = $node->val;
            if($node->right){
                array_push($stack, $node->right);    
            }
            if($node->left){
                array_push($stack, $node->left);    
            }
        }
        return $list;
    }
}
