package main

import "fmt"

type LRUCache struct {
	kvStore map[int]*linkNode
	headNode *linkNode
	tailNode *linkNode
	cap int
	size int
}

type linkNode struct{
	prev *linkNode
	last *linkNode
	key int
	value int
}

func Constructor(capacity int) LRUCache {
	var headNode = newNode()
	var tailNode = newNode()
	headNode.last = tailNode
	tailNode.prev = headNode
	return LRUCache{
		kvStore:  make(map[int]*linkNode),
		headNode: headNode,
		tailNode: tailNode,
		cap:      capacity,
		size:     0,
	}
}

func newNode() *linkNode{
	return &linkNode{
		prev:  nil,
		last:  nil,
		key: 0,
		value: 0,
	}
}

func (this *LRUCache) moveToHead(node *linkNode){
	node.prev.last = node.last
	node.last.prev = node.prev
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *linkNode){
	node.prev = this.headNode
	node.last = this.headNode.last
	this.headNode.last.prev = node
	this.headNode.last = node
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.kvStore[key]; !ok{
		return -1
	}
	this.moveToHead(this.kvStore[key])
	return this.kvStore[key].value
}

func (this *LRUCache) Put(key int, value int)  {
	if _, ok := this.kvStore[key]; !ok{
		var node = newNode()
		node.key = key
		node.value = value
		this.kvStore[key] = node
		this.addToHead(this.kvStore[key])
		this.size++
		if this.size > this.cap{
			node = this.tailNode.prev
			node.prev.last = this.tailNode
			this.tailNode.prev = node.prev
			delete(this.kvStore, node.key)
			this.size--
		}
		return
	}
	this.kvStore[key].value = value
	this.moveToHead(this.kvStore[key])
}

func main(){
	var lRUCache = Constructor(2)
	lRUCache.Put(1, 1)
	lRUCache.Put(2, 2)
	fmt.Println(lRUCache.Get(1))
	lRUCache.Put(3, 3)
	fmt.Println(lRUCache.Get(2))
	lRUCache.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	fmt.Println(lRUCache.Get(1))    // 返回 -1 (未找到)
	fmt.Println(lRUCache.Get(3))    // 返回 3
	fmt.Println(lRUCache.Get(4))    // 返回 4
}