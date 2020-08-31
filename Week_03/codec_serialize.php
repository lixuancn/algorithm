<?php
//二叉树的序列化和反序列化，我把二叉树前序遍历生成一个数组，反序列化就是把这个数组还原成二叉树
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($value) { $this->val = $value; }
 * }
 */

class Codec {
    private $offset = 0;

    function __construct() {

    }

    /**
     * @param TreeNode $root
     * @return String
     */
    function serialize($root) {
        $ret = $this->_serialize($root, array());
        return implode(',', $ret);
    }

    private function _serialize($node, $ret){
        if(!$node){
            $ret[] = "*";
            return $ret;
        }
        $ret[] = $node->val;
        $ret = $this->_serialize($node->left, $ret);
        $ret = $this->_serialize($node->right, $ret);
        return $ret;
    }

    /**
     * @param String $data
     * @return TreeNode
     */
    function deserialize($data) {
        $data = explode(',', $data);
        $data = $this->_deserialize($data, 0);
        return $data;
    }

    private function _deserialize($data, $i) {
        if($data[$i]=="*"){
            return null;
        }
        $node = new TreeNode($data[$i]);
        $node->left = $this->_deserialize($data, ++$this->offset);
        $node->right = $this->_deserialize($data, ++$this->offset);
        return $node;
    }
}

/**
 * Your Codec object will be instantiated and called as such:
 * $obj = Codec();
 * $data = $obj->serialize($root);
 * $ans = $obj->deserialize($data);
 */
$obj = Codec();
$ans = $obj->deserialize([1,2,3,null,null,4,5]);
