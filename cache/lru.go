package cache

type Element struct {
	prev, next *Element
	Key        interface{}
	Value      interface{}
}

func (e *Element) Prev() *Element {
	return e.prev
}

func (e *Element) Next() *Element {
	return e.next
}

type LRUCache struct {
	cache      map[interface{}]*Element
	head, tail *Element
	size       int
}

// 初始化
func New(size int) *LRUCache {
	cache := &LRUCache{
		cache: make(map[interface{}]*Element),
		size:  size,
	}
	return cache
}

func (lr LRUCache) Put(key interface{}, value interface{}) {
	// cache 存在
	if e, ok := lr.cache[key]; ok {
		// cache赋新值
		e.Value = value
		lr.cache[key] = e
		// 移到头结点
		lr.refresh(e)
	}

	if lr.size <= 0 {
		return
	}

	if len(lr.cache) >= lr.size {
		// 没有空间 需删掉最近最少使用的 即tail节点
		lr.remove(lr.tail)
	}
	// 直接存
	e := &Element{Key: key, Value: value}
	lr.cache[key] = e
	if len(lr.cache) == 1 {
		lr.tail = lr.head
	} else if len(lr.cache) == 0 {
		lr.tail = e
	} else {
		lr.head.prev = e
	}
	lr.head = e
}

func (lr LRUCache) Get(key interface{}) *Element {
	// cache 存在
	if e, ok := lr.cache[key]; ok {
		return e
	}
	return nil
}

// refresh 交换节点
func (lr LRUCache) refresh(e *Element) {
	// if tail 节点
	if e.next == nil {
		e.prev.next = nil
	} else {
		e.prev.next = e.next
	}

	// 移到头结点
	e.next = lr.head
	lr.head = e
}

// remove 移除节点
func (lr LRUCache) remove(e *Element) {
	if e.next == nil {
		// tail
		lr.tail = e.prev
		e.prev.next = nil
	} else if e.prev == nil {
		// head
		lr.head = e.next
		e.next = nil
	} else {
		// 中间
		e.prev = e.next
	}
}
