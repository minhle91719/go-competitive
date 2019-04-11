package lru

import (
	"sync"
	"sync/atomic"
	"time"
)

const MaxCacheElement uint64 = 1000000

type Node struct {
	key         string
	value       interface{}
	timeCreated time.Time
	// time add elemet to cache
	next *Node
	prev *Node
}
type LRUCache struct {
	size  uint64
	cache map[string]*Node

	head *Node
	tail *Node

	mu sync.RWMutex
}

func NewNode(key string, value interface{}) *Node {
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
	nNext.prev, n.next = n, nNext
}
func (n *Node) setPrev(nPrev *Node) {
	if nPrev == nil {
		n.prev = nil
		return
	}
	nPrev.next, n.prev = n, nPrev
}
func (n *Node) setValue(value interface{}) {
	n.value = value
	n.timeCreated = time.Now()
}

func NewLRUCache() *LRUCache {
	return &LRUCache{
		size:  0,
		cache: make(map[string]*Node),
		head:  nil,
		tail:  nil,
		mu:    sync.RWMutex{},
	}
}
func (lru *LRUCache) Get(key string) interface{} {
	var value interface{}
	lru.mu.RLock()
	value = lru.cache[key].value
	lru.mu.RUnlock()
	return value
}
func (lru *LRUCache) Set(key string, value interface{}) {
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
	lru.checkSize()
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
