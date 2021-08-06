package _3

import (
	"container/list"
	_ "container/list"
)

/*
1、合并有序链表
2、反转链表
3、单例模式 饿汉模式（一开始就初始化）/懒汉模式需要时初始化
4、简单工厂模式
5、快速排序
6、归并排序
7、实现一个堆排序
8、设计LRU缓存
9、重排链表
10、奇偶链表
11、写三个线程交替打印ABC
12、Top K问题
13、布隆过滤器原理与优点
*/

type AnyType int

type pair struct {
	Key int
	Val AnyType
}

type LRU struct {
	Capacity int
	Cache    map[int]*list.Element
	list     *list.List
}

func NewLRU(cap int) *LRU {
	return &LRU{
		Capacity: cap,
		Cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (l *LRU) Get(key int) AnyType {
	if val, ok := l.Cache[key]; ok {
		l.list.MoveToFront(val)
		return val.Value.(pair).Val
	}
	return AnyType(-1)
}

func (l *LRU) Set(key int, value AnyType) {
	// 值存在 移动到对头并更新val
	// 如果不存在
	// 如果容量充足 直接加入
	// 容量不充足 删除最近最少使用

	if val, ok := l.Cache[key]; ok {
		l.list.MoveToFront(val)
		val.Value = pair{
			Key: key,
			Val: value,
		}
	} else {
		if l.Capacity == len(l.Cache) {
			delete(l.Cache, l.list.Back().Value.(pair).Key)
			l.list.Remove(l.list.Back())
		}
		l.list.PushFront(pair{
			Key: key,
			Val: value,
		})
		l.Cache[key] = l.list.Front()
	}
}
