<?php
//N叉树 - 前序遍历 - 方法一：递归

/**
 * Definition for a Node.
 * class Node {
 *     public $val = null;
 *     public $children = null;
 *     function __construct($val = 0) {
 *         $this->val = $val;
 *         $this->children = array();
 *     }
 * }
 */

class Solution {
    private $list = array();
    /**
     * @param Node $root
     * @return integer[]
     */
    function preorder($root) {
        $this->list[] = $root->val;
        foreach($root->children as $node){
            $this->preorder($node);
        }
        return $this->list;
    }
}
