<?php

/**
 * 双指针，代表两个跑步运动员，一个快跑（一次两步），一个慢跑（一次一步）
 * 如果有环，快跑的总会和慢跑的相遇。
 * Class Solution
 */
class Solution {

    /**
     * @param ListNode $head
     * @return Boolean
     */
    function hasCycle($head) {
        if(!$head || !$head->next){
            return false;
        }
        $slow = $head;
        $fast = $head->next;
        while($fast != $slow){
            if(!$fast || !$fast->next){
                return false;
            }
            $slow = $slow->next;
            $fast = $fast->next->next;
        }
        return true;
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

$head = initListNode(array(3,2,0,-4));
$nodeList = loopListNode($head);
print_r($nodeList);

$obj = new Solution();
$head = $obj->hasCycle($head);
$nodeList = loopListNode($head);
print_r($nodeList);
