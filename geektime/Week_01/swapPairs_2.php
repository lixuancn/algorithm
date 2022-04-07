<?php

/**
 * 遍历的方式两两交换
 * Class Solution
 */
class Solution {

    /**
     * @param ListNode $head
     * @return ListNode
     */
    function swapPairs($head) {
        $dummy = new stdClass();
        $dummy->next = $head;
        $prev = $dummy;
        while($head && $head->next){
            $first = $head;
            $second = $head->next;

            $prev->next = $second;
            $first->next = $second->next;
            $second->next = $first;

            $prev = $first;
            $head = $first->next;
        }
        return $dummy->next;
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
