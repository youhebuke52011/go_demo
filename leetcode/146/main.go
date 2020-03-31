package main

import (
	"fmt"
	"sync"
)

/**
lru缓存实现：哈希表和双向链表,因为要在O（1）时间复杂度删除链表任意一个元素,这道题不用,但如果要手动删除某个缓存或者有过期时间就需要
单链表能实现新来的放头部，最久不用的在尾部删除。但删除的时候需要遍历到尾部，因为单链表只有头指针。在用到已经用到过的数据时，
还要遍历整合链表，来确定是否用过，然后再遍历到响应位置来剔除的节点，并重新放在头部。这效率可想而知。
队列也不行，队列只能做到先进先出，但是重复用到中间的数据时无法把中间的数据移动到顶端。
为什么要用链表，直接用哈希表不行吗？不行，因为要维护次序，双向链表head是最新使用过的，tail是最久没用过的
*/

// 双向链表
type DoubleLinked struct {
	key, value int
	// 过期时间 时间戳,至于啥时候去扫描过期元素,有待商榷
	//Expire int64
	prev, next *DoubleLinked
}

// 哈希表，value是双向链表
type LRUCache struct {
	Capacity   int
	Cache      map[int]*DoubleLinked
	head, tail *DoubleLinked
	// 需要线程安全的话 阔以加互斥锁
	mu sync.RWMutex
	// 过期时间 key:时间戳,至于啥时候去扫描过期元素,阔以参考redis过期策略
	Expire map[int]int64
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Capacity: capacity,
		Cache:    make(map[int]*DoubleLinked, capacity),
	}
	return cache
}

func (this *LRUCache) Get(key int) int {
	this.mu.RLock()
	defer this.mu.RUnlock()
	res := -1
	if node, ok := this.Cache[key]; ok {
		res = node.value
		// 双向链表移到头部
		this.moveNodeToHead(node)
	}
	return res
}

func (this *LRUCache) Put(key int, value int) {
	this.mu.Lock()
	defer this.mu.Unlock()
	if node, ok := this.Cache[key]; ok {
		//this.Cache[key].value = value
		if node.value != value {
			node.value = value
			this.moveNodeToHead(node)
		}
		return
	}
	if this.Capacity <= 0 {
		return
	}

	if len(this.Cache) == this.Capacity {
		// 删掉哈希表中的最近最少使用的
		delete(this.Cache, this.tail.key)
		//	把双线链表末尾元素删掉, 添加到头部
		this.removeTailNode()
	}
	node := &DoubleLinked{
		key:   key,
		value: value,
	}
	this.Cache[key] = node
	//	放到双向链表头部
	this.insertNodeToHead(node)
}

// 把最近使用的移动到链表头部
func (this *LRUCache) moveNodeToHead(node *DoubleLinked) {
	switch node {
	case this.head:
		return
	case this.tail:
		this.removeTailNode()
	default:
		// 删掉node
		node.next.prev = node.prev
		node.prev.next = node.next
	}

	// 再插入到头
	this.insertNodeToHead(node)
}

// 把元素插入到链表头部
func (this *LRUCache) insertNodeToHead(node *DoubleLinked) {
	if this.tail == nil {
		this.tail = node
	} else {
		node.next = this.head
		this.head.prev = node
	}
	this.head = node
}

// 删除链表最后一个元素
func (this *LRUCache) removeTailNode() {
	if this.tail.prev != nil {
		this.tail.prev.next = nil
	} else {
		this.head = nil
	}
	this.tail = this.tail.prev
}

func main() {
	cache := Constructor(2)
	cache.Put(1, 0)
	//fmt.Println(cache.Get(1))
	//cache.Put(1, 2)
	//fmt.Println(cache.Get(1))
	cache.Put(1, 1)
	fmt.Println(cache.Get(1))
	cache.Put(2, 2)
	fmt.Println(cache.Get(2))
	cache.Put(3, 55)
	fmt.Println(cache.Get(3))
}
