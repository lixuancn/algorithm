<?php

//验证有效二叉搜索树

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
     * @return Boolean
     */
    function isValidBST($root)
    {
        if (!$root) {
            return true;
        }
        return $this->handler($root, null, null);
    }

    private function handler($root, $min, $max){
        if(!$root){
            return true;
        }
        if (!is_null($min) && $root->val <= $min){
            return false;
        }
        if (!is_null($max) && $root->val >= $max){
            return false;
        }
        return $this->handler($root->left, $min, $root->val) && $this->handler($root->right, $root->val, $max);
    }
}

