package Emitter

import "sync"

type SyncEmitter struct {
	enabled    bool
	enabledMu  *sync.RWMutex
	handlers   []func(interface{})
	handlersMu *sync.RWMutex
}

func NewSyncEmitter() *SyncEmitter {
	return &SyncEmitter{false,
		new(sync.RWMutex),
		[]func(interface{}){},
		new(sync.RWMutex),
	}
}
func (e *SyncEmitter) Enable() {
	e.enabledMu.Lock()
	defer e.enabledMu.Unlock()
	e.enabled = true
}
func (e *SyncEmitter) Disable() {
	e.enabledMu.Lock()
	defer e.enabledMu.Unlock()
	e.enabled = false
}
func (e *SyncEmitter) AddHandler(handler func(interface{})) {
	e.handlersMu.Lock()
	defer e.handlersMu.Unlock()
	e.handlers = append(e.handlers, handler)
}
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
