<?php

/**
 * 循环的方式翻转链表
 * Class Solution
 */
class Solution {

    /**
     * @param ListNode $head
     * @return ListNode
     */
    function reverseList($head) {
        $cur = &$head;
        $prev = null;
        while($cur != null){
            $tmp = $cur->next;
            $cur->next = $prev;
            $prev = $cur;
            $cur = $tmp;
        }
        return $prev;
    }
}

class ListNode
{
    public $val = 0;
    public $next = null;

    function __construct($val)
    {
        $this->val = $val;
    }
}

function initListNode($list){
    $prev = null;
    $node = null;
    for($i=count($list)-1; $i>=0; $i--){
        $node = new ListNode($list[$i]);
        $node->next = $prev;
        $prev = $node;
    }
    return $node;
}

function loopListNode($head){
    $ret = array();
    $cur = $head;
    while($cur != null){
        $ret[] = $cur->val;
        $cur = $cur->next;
    }
    return $ret;
}

$head = initListNode(array(1,2,3,4,5));
$nodeList = loopListNode($head);
print_r($nodeList);

$obj = new Solution();
$head = $obj->reverseList($head);
$nodeList = loopListNode($head);
print_r($nodeList);
