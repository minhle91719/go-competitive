package cache

import (
	"sync/atomic"
	"time"
)

const MaxCacheElement uint64 = 1000000

type typeData string
type Node struct {
	key         string
	value       typeData
	timeCreated time.Time
	next        *Node
	prev        *Node
}
type LRUCache struct {
	size  uint64
	cache map[string]*Node

	head *Node
	tail *Node
	// time add cache
}

func NewNode(key string, value typeData) *Node {
	n := &Node{
		key:         key,
		value:       value,
		timeCreated: time.Now(),
	}
	return n
}
func (n *Node) setNext(nNext *Node) {
	if nNext == nil {
		n.next = nil
		return
	}
	nNext.prev = n
	n.next = nNext
}
func (n *Node) setPrev(nPrev *Node) {
	if nPrev == nil {
		n.prev = nil
		return
	}
	nPrev.next = n
	n.prev = nPrev
}
func (n *Node) setValue(value typeData) {
	n.value = value
	n.timeCreated = time.Now()
}

func NewLRUCache() *LRUCache {
	return &LRUCache{
		size:  0,
		cache: make(map[string]*Node),
		head:  nil,
		tail:  nil,
	}
}
func (lru *LRUCache) Get(key string) Node {
	return *lru.cache[key]
}
func (lru *LRUCache) Set(key string, value typeData) {
	// check value exist
	var node *Node

	if elm, ok := lru.cache[key]; ok {
		elm.setValue(value)
		node = elm
	} else {
		node = NewNode(key, value)
		lru.cache[key] = node
		atomic.AddUint64(&lru.size, 1)
	}
	lru.movetoEnd(node)
}
func (lru *LRUCache) removeHead() {
	if lru.head != nil {
		delete(lru.cache, lru.head.key)
		lru.head = lru.head.next
	}
}

func (lru *LRUCache) removeLinkedList(n *Node) {
	if n.prev != nil {
		n.prev.setNext(n.next)
	}
	if n.next != nil {
		n.next.setPrev(n.prev)
	}
}
func (lru *LRUCache) movetoEnd(n *Node) {
	lru.removeLinkedList(n)

	if lru.tail == nil { // == head = nil
		lru.head, lru.tail = n, n
	} else {
		n.prev, lru.tail = lru.tail, n
	}
}
func (lru *LRUCache) checkSize() {
	for atomic.LoadUint64(&lru.size) > MaxCacheElement {
		lru.removeHead()
	}
}
