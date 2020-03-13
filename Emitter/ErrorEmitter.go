package Emitter

//This is a implementation of `Emitter`.
//This implementation can only use error as payload.
type ErrorEmitter struct {
	Emitter
}

//Returns a pointer to a asynchronous ErrorEmitter.
func NewAsyncErrorEmitter() *ErrorEmitter {
	return &ErrorEmitter{NewAsyncEmitter()}
}

//Returns a pointer to a synchronous ErrorEmitter.
func NewSyncErrorEmitter() *ErrorEmitter {
	return &ErrorEmitter{NewSyncEmitter()}
}

//Implementation of Emitter.AddHandler.
func (e *ErrorEmitter) AddHandler(handler func(error)) {
	e.Emitter.AddHandler(func(i interface{}) {
		handler(i.(error))
	})
}

//Implementation of Emitter.Emit.
func (e *ErrorEmitter) Emit(err error) {
	e.Emitter.Emit(err)
}
