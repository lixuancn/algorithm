<?php

/**
 * Class Solution
 */
class Solution {

    /**
     * @param ListNode $head
     * @return ListNode
     */
    function detectCycle($head) {
        if(!$head || !$head->next){
            return null;
        }
        $map = array();
        while($head){
            if(!$head || !$head->next){
                return null;
            }
            if(isset($map[$head->val]) && $head == $map[$head->val]){
                return $head;
            }
            $map[$head->val] = $head;
            $head = $head->next;
        }
        return null;
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
