package Emitter

type IndefiniteEmitter interface {
	Enable()
	Disable()
	AddHandler(func(...interface{}))
	Emit(...interface{})
}

type indefinites struct {
	elements []interface{}
}

type SyncIndefiniteEmitter struct {
	*SyncEmitter
}

func NewSyncIndefiniteEmitter() *SyncIndefiniteEmitter {
	return &SyncIndefiniteEmitter{NewSyncEmitter()}
}
func (e *SyncIndefiniteEmitter) AddHandler(handler func(...interface{})) {
	e.SyncEmitter.AddHandler(func(i interface{}) {
		handler(i.(indefinites).elements...)
	})
}
func (e *SyncIndefiniteEmitter) Emit(args ...interface{}) {
	e.SyncEmitter.Emit(indefinites{args})
}

type AsyncIndefiniteEmitter struct {
	*AsyncEmitter
}

func NewAsyncIndefiniteEmitter() *AsyncIndefiniteEmitter {
	return &AsyncIndefiniteEmitter{NewAsyncEmitter()}
}
func (e *AsyncIndefiniteEmitter) AddHandler(handler func(...interface{})) {
	e.AsyncEmitter.AddHandler(func(i interface{}) {
		handler(i.(indefinites).elements...)
	})
}
func (e *AsyncIndefiniteEmitter) Emit(args ...interface{}) {
	e.AsyncEmitter.Emit(indefinites{args})
}
