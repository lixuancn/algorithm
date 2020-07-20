<?php
//翻转一棵树 - 递归 时间复杂度O(n)，因为访问了每个节点一次。
//而因为每个节点都需要被访问一次，所以都时间复杂度最好也就是O(n)了
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
        $left = $this->invertTree($root->left);
        $right = $this->invertTree($root->right);
        $root->left = $right;
        $root->right = $left;
        return $root;
    }
}
