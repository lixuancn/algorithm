<?php
//找树每层最大值的解法和二叉树的层序遍历一毛一样。

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
    function largestValues($root) {
        if(!$root){
            return [];
        }
        $queue = array();
        array_push($queue, $root);
        $ret = array();
        while($queue){
            $l = count($queue);
            $max = null;
            for($i=0; $i<$l; $i++){
                $node = array_shift($queue);
                if(is_null($max) || $max < $node->val){
                    $max = $node->val;
                }
                if($node->left){
                    array_push($queue, $node->left);
                }
                if($node->right){
                    array_push($queue, $node->right);
                }
            }
            $ret[] = $max;

        }
        return $ret;
    }
}
