<?php

//验证有效二叉搜索树 - 中序遍历后是从小到大的

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

    private $ret = array();
    /**
     * @param TreeNode $root
     * @return Boolean
     */
    function isValidBST($root)
    {
        if (!$root) {
            return true;
        }
        $this->inOrder($root);
        if(count($this->ret) == 0){
            return false;
        }
        $preVal = null;
        foreach($this->ret as $val){
            if(is_null($preVal)){
                $preVal = $val;
                continue;
            }
            if($val <= $preVal){
                return false;
            }
            $preVal = $val;
        }
        return true;
    }

    private function inOrder($root){
        if($root->left){
            $this->inOrder($root->left);
        }
        $this->ret[] = $root->val;
        if($root->right){
            $this->inOrder($root->right);
        }
    }
}

