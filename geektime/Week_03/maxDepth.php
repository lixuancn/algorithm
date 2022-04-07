<?php
//二叉树的最大深度
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
    function maxDepth($root) {
        $depth = $this->handler($root, 0);
        return $depth;
    }

    function handler($node, $depth){
        if(!$node){
            return $depth;
        }
        $d1 = $this->handler($node->left, $depth+1);
        $d2 = $this->handler($node->right, $depth+1);
        return max($d1, $d2);
    }
}
