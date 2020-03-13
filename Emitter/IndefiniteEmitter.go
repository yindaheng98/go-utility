package Emitter

//This is the basic interface of multi-payload emitter.
//This type of emitter can carry more than 1 payload at a event.
type IndefiniteEmitter interface {

	//Enable the emitter.
	Enable()

	//Disable the emitter.
	Disable()

	//Add a handler to the emitter.
	//
	//Handler is a function.
	//When `Emit(...)` is called, all the handler added by `AddHandler(func(...))` will be called,
	//and `Emit(...)`'s argument will become handler's argument.
	AddHandler(func(...interface{}))

	//Receive payload and emit a event.
	Emit(...interface{})
}

type indefinites struct {
	elements []interface{}
}

//This is a implementation of IndefiniteEmitter.
//This implementation will handle the events synchronously.
type SyncIndefiniteEmitter struct {
	*SyncEmitter
}

//Returns a pointer to a synchronous IndefiniteEmitter.
func NewSyncIndefiniteEmitter() *SyncIndefiniteEmitter {
	return &SyncIndefiniteEmitter{NewSyncEmitter()}
}

//Implementation of IndefiniteEmitter.AddHandler
func (e *SyncIndefiniteEmitter) AddHandler(handler func(...interface{})) {
	e.SyncEmitter.AddHandler(func(i interface{}) {
		handler(i.(indefinites).elements...)
	})
}

//Implementation of IndefiniteEmitter.Emit
func (e *SyncIndefiniteEmitter) Emit(args ...interface{}) {
	e.SyncEmitter.Emit(indefinites{args})
}

//This is a implementation of IndefiniteEmitter.
//This implementation will handle the events asynchronously.
type AsyncIndefiniteEmitter struct {
	*AsyncEmitter
}

//Returns a pointer to a asynchronous IndefiniteEmitter.
func NewAsyncIndefiniteEmitter() *AsyncIndefiniteEmitter {
	return &AsyncIndefiniteEmitter{NewAsyncEmitter()}
}

//Implementation of IndefiniteEmitter.AddHandler
func (e *AsyncIndefiniteEmitter) AddHandler(handler func(...interface{})) {
	e.AsyncEmitter.AddHandler(func(i interface{}) {
		handler(i.(indefinites).elements...)
	})
}

//Implementation of IndefiniteEmitter.Emit
func (e *AsyncIndefiniteEmitter) Emit(args ...interface{}) {
	e.AsyncEmitter.Emit(indefinites{args})
}
