package channel

type Fun func(key string) (interface{}, error)

type result struct {
	val interface{}
	err error
}

type entry struct {
	res   result
	ready chan struct{}
}

type request struct {
	key      string
	response chan<- result
}

type Memo struct {
	requests chan request
}

func New(f Fun) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (m *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	m.requests <- request{key: key, response: response}
	res := <-response
	return res.val, res.err
}

func (m *Memo) Close() {
	close(m.requests)
}

func (m *Memo) server(f Fun) {
	cache := make(map[string]*entry)
	for req := range m.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{
				ready: make(chan struct{}),
			}
			cache[req.key] = e
			// 通知执行
			go e.call(f, req.key)
		}
		// 投递结果
		go e.delivery(req.response)
	}
}

func (e *entry) call(f Fun, key string) {
	e.res.val, e.res.err = f(key)
	// close掉后<-e.ready不再阻塞
	close(e.ready)
}

func (e *entry) delivery(resp chan<- result) {
	// 准备就绪
	<-e.ready
	// Send the result to the client.
	resp <- e.res
}
