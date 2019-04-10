package hashTable

// hashtable的golang实现版本, 对外api都提供了线程安全锁

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

type IHashTable interface {
	// 添加
	Put(key interface{}, value interface{}) error
	// 获取key对应的值
	Get(key interface{}) interface{}
	// 删除
	Remove(key interface{}) error
	// 当前hashtable的数据量
	Size() int
	// 是否为空
	IsEmpty() bool
	// 打印每一个hash槽内的内容
	Show()
}

type Options struct {
	// hashtable容量，设置默认桶容量
	Capacity uint
	// 负载因子 0<= x <=1
	LoadFactor float64
	// 是否记录扩容log
	Debug bool
}

type Entry struct {
	key   interface{}
	value interface{}
	next  *Entry
}

type HashTable struct {
	// hash桶
	// Hashtable保存key-value的数组。
	// Hashtable是采用拉链法实现的，每一个Entry本质上是一个单向链表
	table []*Entry

	// hashtable的设定容量
	capacity int

	// 加载因子, 超过此比例就自动扩容
	loadFactor float64

	// debug
	debug bool

	// 读写锁
	lock sync.RWMutex

	// 阈值，用于判断是否需要调整Hashtable的容量（threshold = 容量*加载因子）
	threshold int

	// Hashtable中元素的实际数量
	count int
}

var _ IHashTable = &HashTable{}

func NewHashTable(o *Options) *HashTable {
	var caps = o.Capacity
	if caps == 0 {
		caps = 1
	}

	var ht = &HashTable{
		table:      make([]*Entry, caps),
		capacity:   int(o.Capacity),
		loadFactor: o.LoadFactor,
		debug:      o.Debug,
		lock:       sync.RWMutex{},
		threshold:  int(o.LoadFactor * float64(caps)),
	}
	return ht
}

func (h *HashTable) Put(key interface{}, value interface{}) error {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.insert(key, value)

	h.count++

	// 检查是否达到了设定的负载值
	if int(h.count) >= h.threshold {
		h.reHash()
	}

	return nil
}

func (h *HashTable) insert(key interface{}, value interface{}) error {

	findNode := h.getEntry(key)
	if findNode != nil {
		findNode.value = value
		return nil
	}

	newEntry := &Entry{key: key, value: value}

	// 找出key所在的table
	hash := h.hash(key)
	var entry = h.table[hash]

	if entry == nil {
		h.table[hash] = newEntry
	} else {
		for entry.next != nil {
			entry = entry.next
		}
		entry.next = newEntry
	}

	return nil
}

func (h *HashTable) reHash() error {

	var oldTable = h.table
	var oldCap = h.capacity
	// 新的容量为
	var newCap = oldCap*2 + 1

	// 设置新的门槛
	h.threshold = int(float64(newCap) * h.loadFactor)
	// 设置新的容量
	h.capacity = newCap
	// 设置新的entry
	h.table = make([]*Entry, newCap)

	// 记录log
	if h.debug {
		log.Printf("数量到达: %v(%v*%v)，开始扩容: %v -> %v \n",
			h.count, oldCap, h.loadFactor, oldCap, newCap)
	}

	// 将旧的entry复制到新entry
	for i := oldCap - 1; i >= 0; i-- {
		current := oldTable[i]
		for current != nil {
			h.insert(current.key, current.value)
			current = current.next
		}
	}
	return nil
}

func (h *HashTable) Get(key interface{}) interface{} {
	h.lock.RLock()
	defer h.lock.RUnlock()

	e := h.getEntry(key)

	if e != nil {
		return e.value
	}

	return nil
}

func (h *HashTable) getEntry(key interface{}) *Entry {
	// 获取key的hash所在table
	hash := h.hash(key)
	// 获取对应的entry
	entry := h.table[hash]

	if entry == nil {
		return nil
	}

	for entry.key != key && entry.next != nil {
		entry = entry.next
	}

	if entry.key == key {
		return entry
	}

	return nil
}

func (h *HashTable) Remove(key interface{}) error {
	h.lock.Lock()
	defer h.lock.Unlock()

	// 获取key的hash所在table
	hash := h.hash(key)
	// 获取对应的entry
	var prev = h.table[hash]

	if prev == nil {
		return errors.New("the given key is not exists")
	}

	if prev.key == key {
		h.table[hash] = prev.next
		//prev = nil
		return nil
	}

	for prev.key != key && prev.next != nil {
		if prev.next.key == key {
			prev.next = prev.next.next
			//prev = nil
			return nil
		}
		prev = prev.next
	}
	return errors.New("the given key is not exists")
}

func (h *HashTable) Size() int {
	h.lock.RLock()
	defer h.lock.RUnlock()
	return h.count
}

func (h *HashTable) IsEmpty() bool {
	h.lock.RLock()
	defer h.lock.RUnlock()
	return h.count == 0
}

func (h *HashTable) hash(key interface{}) int {
	return int(uint(HashCode(key)) % uint(h.capacity))
}

func (h *HashTable) Show() {
	for k, item := range h.table {
		fmt.Printf("%v: ", k)
		for item != nil {
			fmt.Printf("%v=>%v, ", item.key, item.value)
			item = item.next
		}
		fmt.Println("")
	}
}

func HashCode(key interface{}) int {
	keyStr := fmt.Sprintf("%s", key)
	keyLen := len(keyStr)

	var h = 0

	if keyLen > 0 {
		for i := 0; i < keyLen; i++ {
			h = h<<5 - h + int(keyStr[i])
		}
	}

	return h
}
