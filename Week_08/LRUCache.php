<?php
//双向链表 + hash表
class DLinkedNode{
    public $key;
    public $value;
    public $prev = null;
    public $next = null;
    public function __construct($k, $v)
    {
        $this->key = $k;
        $this->value = $v;
    }
}

class LRUCache{
    public $cache = array();
    private $cap = 0;
    private $size = 0;
    private $head = null;
    private $tail = null;
    public function __construct($cap)
    {
        $this->cap = $cap;
        //两个虚拟节点
        $this->head = new DLinkedNode(null, null);
        $this->tail = new DLinkedNode(null, null);
        $this->head->next = $this->tail;
        $this->tail->prev = $this->head;
    }

    public function get($key){
        if(!isset($this->cache[$key])){
            return -1;
        }
        $node = $this->cache[$key];
        $this->moveToHead($node);
        return $node->value;
    }

    public function put($key, $value){
        if(isset($this->cache[$key])){
            $node = $this->cache[$key];
            $node->value = $value;
            $this->moveToHead($node);
            return;
        }
        $node = new DLinkedNode($key, $value);
        $this->cache[$key] = $node;
        $this->addToHead($node);
        $this->size++;
        if($this->size > $this->cap){
            $removeNode = $this->removeTail();
            unset($this->cache[$removeNode->key]);
            $this->size--;
        }
    }

    private function addToHead($node)
    {
        $node->prev = $this->head;
        $node->next = $this->head->next;
        $this->head->next->prev = $node;
        $this->head->next = $node;
    }

    private function removeNode($node)
    {
        $node->prev->next = $node->next;
        $node->next->prev = $node->prev;
    }

    private function moveToHead($node)
    {
        $this->removeNode($node);
        $this->addToHead($node);
    }

    private function removeTail()
    {
        $node = $this->tail->prev;
        $this->removeNode($node);
        return $node;
    }
}

$obj = new LRUCache(2);
var_dump($obj->put(1, 1));
var_dump($obj->put(2, 2));
var_dump($obj->get(1));       // 返回  1
var_dump($obj->put(3, 3));    // 该操作会使得关键字 2 作废
var_dump("---");
var_dump($obj->get(2));       // 返回 -1 (未找到)
var_dump($obj->put(4, 4));    // 该操作会使得关键字 1 作废
var_dump($obj->get(1));       // 返回 -1 (未找到)
var_dump($obj->get(3));       // 返回  3
var_dump($obj->get(4));       // 返回  4
