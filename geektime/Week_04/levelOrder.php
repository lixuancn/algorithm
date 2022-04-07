<?php
//广度优先，难点在如果记录层，加一层for是很巧妙的办法

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
     * @return Integer[][]
     */
    function levelOrder($root) {
        if(!$root){
            return [];
        }
        $queue = array();
        array_push($queue, $root);
        $ret = array();
        while($queue){
            $l = count($queue);
            $ans = array();
            for($i=0; $i<$l; $i++){
                $node = array_shift($queue);
                $ans[] = $node->val;
                if($node->left){
                    array_push($queue, $node->left);
                }
                if($node->right){
                    array_push($queue, $node->right);
                }
            }
            $ret[count($ret)] = $ans;

        }
        return $ret;
    }
}
