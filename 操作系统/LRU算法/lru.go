package main

import (
	"container/list"
	"fmt"
)

/**
* @Author : awen
* @Date : 2020/2/23 10:13 下午
 */

// 参考： https://zhuanlan.zhihu.com/p/34989978

type LRUCache struct {
	capacity int                   // 队列容量,也是缓存容量
	cache    map[int]*list.Element // 缓存，缓存队列中的元素
	list     *list.List            //链表，模仿队列
}

// 键值对
type Pair struct {
	key   int
	value int
}

//构造器
func Constructor(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

// 获取键值对的值
func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok { // 元素是否存在
		this.list.MoveToFront(elem)    // 最近使用了，处于最活跃状态，移动元素到链表头部
		return elem.Value.(Pair).value // 返回键值对的值，取出来的是interface ，需要转成 Pair
	}
	return -1
}

// 往缓存里设值
func (this *LRUCache) Put(key int, val int) {
	if elem, ok := this.cache[key]; ok { // 缓存中键值对存在
		// 存在
		this.list.MoveToFront(elem) // 把最近使用的元素移到队首
		elem.Value = Pair{key, val} // 返回
	} else { // 键值对不存在
		if this.list.Len() >= this.capacity {
			// 容量已满
			delete(this.cache, this.list.Back().Value.(Pair).key) // 删除缓存中的键值对
			this.list.Remove(this.list.Back()) // 删除队列中最近最久没有使用的元素
		}
		// 否则，直接插入在队列首
		this.list.PushFront(Pair{key, val})
		this.cache[key] = this.list.Front()
	}
}

// 遍历list，打印键值对
func (lruCache *LRUCache) print() {
	for elem := lruCache.list.Front(); elem != nil ; elem = elem.Next() {
		fmt.Printf("key: %d, vlaue:%d \n",elem.Value.(Pair).key,elem.Value.(Pair).value)
	}
}

func main()  {
	lruCache := Constructor(3)
	lruCache.Put(1,1)
	lruCache.Put(2,2)
	lruCache.Put(3,3)
	lruCache.Put(4,4)
	lruCache.Get(3)
	lruCache.Get(2)
	lruCache.print() // 4 3 2
}
