package Emitter

import "sync"

//This is a implementation of Emitter.
//This implementation will handle the events synchronously.
type SyncEmitter struct {
	enabled    bool
	enabledMu  *sync.RWMutex
	handlers   []func(interface{})
	handlersMu *sync.RWMutex
}

//Returns a pointer to a synchronous Emitter.
func NewSyncEmitter() *SyncEmitter {
	return &SyncEmitter{false,
		new(sync.RWMutex),
		[]func(interface{}){},
		new(sync.RWMutex),
	}
}

//Implementation of Emitter.Enable.
func (e *SyncEmitter) Enable() {
	e.enabledMu.Lock()
	defer e.enabledMu.Unlock()
	e.enabled = true
}

//Implementation of Emitter.Disable.
func (e *SyncEmitter) Disable() {
	e.enabledMu.Lock()
	defer e.enabledMu.Unlock()
	e.enabled = false
}

//Implementation of Emitter.AddHandler.
func (e *SyncEmitter) AddHandler(handler func(interface{})) {
	e.handlersMu.Lock()
	defer e.handlersMu.Unlock()
	e.handlers = append(e.handlers, handler)
}

//Implementation of Emitter.Emit.
func (e *SyncEmitter) Emit(info interface{}) {
	e.enabledMu.RLock()
	e.handlersMu.RLock()
	defer e.handlersMu.RUnlock()
	defer e.enabledMu.RUnlock()
	if e.enabled {
		for _, handler := range e.handlers {
			handler(info)
		}
	}
}
