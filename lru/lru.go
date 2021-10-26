package lru

import (
	"container/list"
	"sync"
)

var _ = list.List{}

type Entry struct {
	Key   int
	Value int
}

type LRU struct {
	UseList  *list.List
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
		UseList:  list.New(),
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
		l.UseList.MoveToFront(val)
		val.Value = &Entry{
			Key:   key,
			Value: v,
		}
		return
	}
	elm := l.UseList.PushFront(&Entry{
		Key:   key,
		Value: v,
	})
	l.VMap[key] = elm
	// 超出限制
	if l.UseList.Len() > l.Capacity {
		// 曾经出错的地方
		delete(l.VMap, l.UseList.Back().Value.(*Entry).Key)
		l.UseList.Remove(l.UseList.Back())
	}
}

func (l *LRU) Get(key int) (int, bool) {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()
	l.GetC++
	if val, ok := l.VMap[key]; ok {
		l.HitC++
		l.UseList.MoveToFront(val)
		return val.Value.(*Entry).Value, true
	}
	return -1, false
}
