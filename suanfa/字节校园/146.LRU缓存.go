package main

import (
	"math"
)

type LRUCache struct {
	Capacity int
	CahceMap map[int]int
	TimesMap map[int]int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		CahceMap: make(map[int]int, capacity),
		TimesMap: make(map[int]int, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if this.CahceMap[key] != 0 {
		this.TimesMap[key] += 1
		return this.CahceMap[key]
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.CahceMap[key] != 0 {
		this.CahceMap[key] = value
		this.TimesMap[key] += 1
	} else {
		if len(this.CahceMap) == this.Capacity {
			minTimesKey := math.MaxInt
			minTimes := math.MaxInt
			for i, v := range this.TimesMap {
				if v < minTimes {
					minTimes = v
					minTimesKey = i
				}
			}
			delete(this.CahceMap, minTimesKey)
			delete(this.TimesMap, minTimesKey)
		}
		this.CahceMap[key] = value
		this.TimesMap[key] = 0
	}

}

func main() {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1) // 缓存是 {1=1}
	lRUCache.Put(2, 2) // 缓存是 {1=1, 2=2}
	lRUCache.Get(1)    // 返回 1
	lRUCache.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	lRUCache.Get(2)    // 返回 -1 (未找到)
	lRUCache.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	lRUCache.Get(1)    // 返回 -1 (未找到)
	lRUCache.Get(3)    // 返回 3
	lRUCache.Get(4)    // 返回 4

}
