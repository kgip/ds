package ds

import "sync"

type CacheNode struct {
	Key  string
	Val  interface{}
	Pre  *CacheNode
	Next *CacheNode
}

type LRUCache struct {
	cache    map[string]*CacheNode
	head     *CacheNode
	tail     *CacheNode
	lock     *sync.RWMutex
	MaxCount int //最大元素数量，超过这个数量时再添加元素则要淘汰已有元素
}

func (cache *LRUCache) Put(key string, val interface{}) {
	cache.lock.Lock()
	defer cache.lock.Unlock()
	//已经存在这个key，更新元素
	if v, ok := cache.cache[key]; ok {
		v.Val = val
		if len(cache.cache) > 1 {
			if cache.head == v {
				cache.head = cache.head.Next
				cache.head.Pre = nil
			} else {
				v.Pre.Next = v.Next
				v.Next.Pre = v.Pre
			}
			cache.tail.Next = v
			v.Next = nil
			v.Pre = cache.tail
			cache.tail = v
		}
	} else {                                    //添加一个新key
		if cache.MaxCount <= len(cache.cache) { //删除老key
			delete(cache.cache, cache.head.Key)
			cache.head = cache.head.Next
			cache.head.Pre = nil
		}
		//创建新key
		node := &CacheNode{Key: key, Val: val}
		cache.cache[key] = node
		if cache.head == nil { //添加第一个元素
			cache.head = node
			cache.tail = node
		} else {
			cache.tail.Next = node
			node.Next = nil
			node.Pre = cache.tail
			cache.tail = node
		}
	}
}

func (cache *LRUCache) Get(key string) (val interface{}, exist bool) {
	cache.lock.Lock()
	defer cache.lock.Unlock()
	if v, ok := cache.cache[key]; ok {
		if len(cache.cache) > 1 {
			if cache.head == v {
				cache.head = cache.head.Next
				cache.head.Pre = nil
			} else {
				v.Pre.Next = v.Next
				v.Next.Pre = v.Pre
			}
			cache.tail.Next = v
			v.Next = nil
			v.Pre = cache.tail
			cache.tail = v
		}
		return v.Val, true
	}
	return nil, false
}

func (cache *LRUCache) Delete(key string) {
	cache.lock.Lock()
	defer cache.lock.Unlock()
	if v, ok := cache.cache[key]; ok {
		delete(cache.cache, key)
		if len(cache.cache) <= 1 {
			cache.head = nil
			cache.tail = nil
		} else {
			if v == cache.head {
				cache.head = cache.head.Next
				cache.head.Pre = nil
			} else if v == cache.tail {
				cache.tail = cache.tail.Pre
				cache.tail.Next = nil
			} else {
				v.Pre.Next = v.Next
				v.Next.Pre = v.Pre
				v.Next = nil
				v.Pre = nil
			}
		}
	}
}
