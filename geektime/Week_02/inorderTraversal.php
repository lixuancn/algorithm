<?php
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
    function inorderTraversal($root) {
        if($root->left){
            $this->inorderTraversal($root->left);
        }
        $this->list[] = $root->val;
        if($root->right){
            $this->inorderTraversal($root->right);
        }
        return $this->list;
    }
}
