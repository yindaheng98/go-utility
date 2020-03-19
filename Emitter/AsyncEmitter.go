package Emitter

import (
	"github.com/yindaheng98/go-utility/Single"
	"sync"
)

//This is a implementation of Emitter.
//This implementation will handle the events asynchronously.
type AsyncEmitter struct {
	runner     *Single.Processor    //控制事件处理线程
	handlers   *[]func(interface{}) //事件处理器列表
	handlersMu *sync.RWMutex        //事件处理器列表读写锁
	events     chan interface{}     //事件队列
	eventsMu   *sync.RWMutex        //事件队列的新建删除和使用操作锁
	enabled    bool                 //启停标记
	enabledMu  *sync.RWMutex        //启停标记读写锁
}

//Returns a pointer to a asynchronous Emitter.
func NewAsyncEmitter() *AsyncEmitter {
	e := &AsyncEmitter{Single.NewProcessor(),
		new([]func(interface{})), new(sync.RWMutex),
		make(chan interface{}, 10), new(sync.RWMutex),
		false, new(sync.RWMutex)}
	e.runner.Callback.Started = func() {
		e.enabledMu.Lock()
		defer e.enabledMu.Unlock()
		e.enabled = true
	}
	e.runner.Callback.Stopped = func() {
		e.enabledMu.Lock()
		defer e.enabledMu.Unlock()
		e.enabled = false
	}
	return e
}

//Implementation of Emitter.AddHandler.
func (e *AsyncEmitter) AddHandler(handler func(interface{})) {
	e.handlersMu.Lock()
	defer e.handlersMu.Unlock()
	*e.handlers = append(*e.handlers, handler)
}

//Implementation of Emitter.Emit.
func (e *AsyncEmitter) Emit(info interface{}) {
	defer func() {
		if recover() != nil {
			e.Disable()
		}
	}()
	e.enabledMu.RLock()
	defer e.enabledMu.RUnlock()
	if e.enabled { //只有不在disabled状态才入队列
		e.eventsMu.RLock()
		defer e.eventsMu.RUnlock()
		e.events <- info
	}
}

//Implementation of Emitter.Enable.
func (e *AsyncEmitter) Enable() {
	e.runner.Start(e.eventLoop)
}

//Implementation of Emitter.Disable.
func (e *AsyncEmitter) Disable() {
	e.runner.Stop()
	e.eventsMu.Lock()
	defer e.eventsMu.Unlock()
	close(e.events)
	e.events = make(chan interface{})
}

//事件处理循环：出队列处理事件
func (e *AsyncEmitter) eventLoop() {
	info, ok := <-e.events
	if !ok {
		return
	}
	e.handlersMu.RLock()
	defer e.handlersMu.RUnlock()
	for _, handler := range *e.handlers {
		handler(info)
	}
}
