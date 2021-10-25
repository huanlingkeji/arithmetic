package lru

import (
	"container/list"
	"sync"
)

var _ = list.List{}

type LRU struct {
	UseList  list.List
	VMap     map[int]*list.Element
	Capacity int
	GetC     int
	HitC     int
	Mutex    sync.Mutex
}

type Node struct {
	V int
}

func NewLRU(cap int) *LRU {
	return &LRU{
		UseList:  list.List{},
		VMap:     make(map[int]*list.Element),
		Capacity: cap,
		GetC:     0,
		HitC:     0,
		Mutex:    sync.Mutex{},
	}
}

func (l *LRU) Put(key, v int) {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()
	// key存在
	if val, ok := l.VMap[key]; ok {
		l.UseList.PushFront(val)
		val.Value = v
		return
	}
	l.UseList.PushFront(v)
	l.VMap[v] = l.UseList.Front()
	// 超出限制
	if len(l.VMap) >= l.Capacity {
		// key不存在
		l.UseList.Remove(l.UseList.Back())
	}
}

func (l *LRU) Get(key int) (interface{}, bool) {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()
	l.GetC++
	if val, ok := l.VMap[key]; ok {
		l.HitC++
		l.UseList.MoveToFront(val)
		return val.Value, true
	}
	return nil, false
}
