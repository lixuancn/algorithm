<?php

/**
 * 递归的方式两两交换
 * 思想是：第一个节点的next是第二个节点的next，第二个节点的next是第一个节点。
 * Class Solution
 */
class Solution {

    /**
     * @param ListNode $head
     * @return ListNode
     */
    function swapPairs($head) {
        if(!$head || !$head->next){
            return $head;
        }
        $first = $head;
        $second = $head->next;
        $first->next = $this->swapPairs($second->next);
        $second->next = $first;
        return $second;
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

$head = initListNode(array(1,2,3,4));
$nodeList = loopListNode($head);
print_r($nodeList);

$obj = new Solution();
$head = $obj->swapPairs($head);
$nodeList = loopListNode($head);
print_r($nodeList);
