<?php

//二叉树 - 最小深度
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
     * @return Integer
     */
    function minDepth($root) {
        if(!$root){
            return 0;
        }
        if(!$root->left && !$root->right){
            return 1;
        }
        $minDepth = pow(2, 31);
        if($root->left){
            $minDepth = min($this->minDepth($root->left), $minDepth);
        }
        if($root->right){
            $minDepth = min($this->minDepth($root->right), $minDepth);
        }
        return $minDepth+1;
    }
}

