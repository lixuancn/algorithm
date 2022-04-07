<?php
//递归
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
    function rob($root) {
        $result = $this->dfs($root);
        //0是偷，1是不偷
        return max($result[0], $result[1]);
    }

    function dfs($node){
        if(!$node){
            return array(0, 0);
        }
        $l = $this->dfs($node->left);
        $r = $this->dfs($node->right);
        $result = array();
        $result[0] = $node->val + $l[1] + $r[1];
        $result[1] = max($l[0], $l[1]) + max($r[0], $r[1]);
        return $result;
    }
}
