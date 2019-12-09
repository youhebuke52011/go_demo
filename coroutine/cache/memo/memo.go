package memo

import (
	"sync"
)

type Fun func(key string) (interface{}, error)

type result struct {
	val interface{}
	err error
}

type entry struct {
	res   result
	ready chan struct{}
}

type Memo struct {
	// 需要缓存的函数
	f Fun
	// 需要缓存的值
	cache map[string]*entry
	// 加锁
	mx sync.Mutex
}

func New(f Fun) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]*entry),
	}
}

var flag int
//
func (m *Memo) Get(key string) (interface{}, error) {
	// 获取互斥锁来保护共享变量cache map
	m.mx.Lock()
	e := m.cache[key]
	flag += 1
	if e == nil {
		// 初始化entry
		e := &entry{
			ready: make(chan struct{}),
		}
		m.cache[key] = e
		m.mx.Unlock()

		// 请求函数,填充 value
		e.res.val, e.res.err = m.f(key)

		// 广播准备条件
		close(e.ready)
	} else {
		m.mx.Unlock()
		//  wait for ready condition 读取操作在channel关闭之前一直是阻塞
		<-e.ready
	}


	if e == nil {
		return []byte(string(flag)), nil
	}
	return e.res.val, e.res.err
}

// not safee
//func (m *Memo) Get(key string) (interface{}, error) {
//	res, ok := m.cache[key]
//	if !ok {
//		res.val, res.err = m.f(key)
//		m.cache[key] = res
//	}
//	return res.val, res.err
//}

// 加互斥锁  效率极低  犹如串行
//func (m *Memo) Get(key string) (interface{}, error) {
//	//m.mx.Lock()
//	//defer m.mx.Unlock()
//	res, ok := m.cache[key]
//	//m.mx.Unlock()
//	if !ok {
//		res.val, res.err = m.f(key)
//		//m.mx.Lock()
//		m.cache[key] = res
//		//m.mx.Unlock()
//	}
//	return res.val, res.err
//}
