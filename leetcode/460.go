package main

import "fmt"

type LFUCache struct {
	//某频次下的头结点，就一个节点
	recent map[int]*Node
	//某频次下有多少个节点
	count map[int]int
	//hashmap（和LRU一样）
	cache map[int]*Node
	//双向链表（和LRU一样）
	list     *DoublyList
	capacity int
}

type DoublyList struct {
	head *Node
	tail *Node
}

type Node struct {
	prev      *Node
	next      *Node
	key       int
	val       int
	frequency int
}

func newDoublyList() *DoublyList {
	doublyList := &DoublyList{
		head: &Node{},
		tail: &Node{},
	}
	doublyList.head.next = doublyList.tail
	doublyList.tail.prev = doublyList.head
	return doublyList
}

//从链表中删除
func removeNode(node *Node) {
	node.next.prev = node.prev
	node.prev.next = node.next
}

//插入到链表头上
func addBefore(curNode, newNode *Node) {
	curNode.prev.next = newNode
	newNode.prev = curNode.prev
	newNode.next = curNode
	curNode.prev = newNode

	//pre := currNode.prev
	//pre.next = newNode
	//newNode.next = currNode
	//currNode.prev = newNode
	//newNode.prev = pre
}

func (dl *DoublyList) removeTail() {
	dl.tail.prev.prev.next = dl.tail
	dl.tail.prev = dl.tail.prev.prev

	//pre := dl.tail.prev.prev
	//pre.next = dl.tail
	//dl.tail.prev = pre
}

func (dl *DoublyList) isEmpty() bool {
	return dl.head.next == dl.tail
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		recent:   make(map[int]*Node, capacity),
		count:    make(map[int]int),
		cache:    make(map[int]*Node, capacity),
		list:     newDoublyList(),
		capacity: capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if this.capacity == 0 {
		return -1
	}
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	//因为再挪动node后，node.next就变了
	next := node.next
	if this.count[node.frequency+1] > 0 {
		//如果频次n+1已经有节点了，那么就把node插入到这个节点的前面，即n+1节点集合的第一位
		removeNode(node)
		addBefore(this.recent[node.frequency+1], node)
	} else if this.count[node.frequency] > 1 && this.recent[node.frequency] != node {
		//如果频次n+1的集合没有节点，频次n的集合有除自己外的其他节点，则把自己插入到频次n的集合第一位（如果已经再第一位则不做操作，就是&&后面的判断）
		removeNode(node)
		addBefore(this.recent[node.frequency], node)
	}
	//更新recent
	this.recent[node.frequency+1] = node
	if this.count[node.frequency] <= 1 {
		this.recent[node.frequency] = nil
	} else if this.recent[node.frequency] == node {
		this.recent[node.frequency] = next
	}
	//更新count
	this.count[node.frequency+1]++
	this.count[node.frequency]--
	node.frequency++
	return node.val
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	node, ok := this.cache[key]
	if ok {
		this.Get(key)
		node.val = value
		return
	}
	//淘汰
	if this.capacity <= len(this.cache) {
		tailNode := this.list.tail.prev
		this.list.removeTail()
		if this.count[tailNode.frequency] <= 1 {
			this.recent[tailNode.frequency] = nil
		}
		this.count[tailNode.frequency]--
		delete(this.cache, tailNode.key)
	}
	node = &Node{
		key:       key,
		val:       value,
		frequency: 1,
	}
	//插入到链表中
	if this.count[1] > 0 {
		addBefore(this.recent[1], node)
	} else {
		addBefore(this.list.tail, node)
	}
	this.cache[key] = node
	this.recent[1] = node
	this.count[1]++
}

func main() {
	lfu := Constructor(2)
	lfu.Put(1, 1)           // cache=[1,_], cnt(1)=1
	lfu.Put(2, 2)           // cache=[2,1], cnt(2)=1, cnt(1)=1
	fmt.Println(lfu.Get(1)) // 返回 1
	// cache=[1,2], cnt(2)=1, cnt(1)=2
	lfu.Put(3, 3) // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小
	// cache=[3,1], cnt(3)=1, cnt(1)=2
	fmt.Println(lfu.Get(2)) // 返回 -1（未找到）
	fmt.Println(lfu.Get(3)) // 返回 3
	// cache=[3,1], cnt(3)=2, cnt(1)=2
	lfu.Put(4, 4) // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用
	// cache=[4,3], cnt(4)=1, cnt(3)=2
	fmt.Println(lfu.Get(1)) // 返回 -1（未找到）
	fmt.Println(lfu.Get(3)) // 返回 3
	// cache=[3,4], cnt(4)=1, cnt(3)=3
	fmt.Println(lfu.Get(4)) // 返回 4
	// cache=[3,4], cnt(4)=2, cnt(3)=3
}
