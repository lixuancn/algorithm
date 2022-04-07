<?php
//前序遍历 - 方法一：递归

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
    private $list = array();
    /**
     * @param TreeNode $root
     * @return Integer[]
     */
    function preorderTraversal($root) {
        $this->list[] = $root->val;
        if($root->left){
            $this->preorderTraversal($root->left);
        }
        if($root->right){
            $this->preorderTraversal($root->right);
        }
        return $this->list;
    }
}
