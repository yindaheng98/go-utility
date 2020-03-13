package Emitter

//This is a implementation of `Emitter`.
//This implementation can only use an error and an interface{} as payload.
type ErrorInfoEmitter struct {
	IndefiniteEmitter
}

//Returns a pointer to a asynchronous ErrorInfoEmitter.
func NewAsyncErrorInfoEmitter() *ErrorInfoEmitter {
	return &ErrorInfoEmitter{NewAsyncIndefiniteEmitter()}
}

//Returns a pointer to a synchronous ErrorInfoEmitter.
func NewSyncErrorInfoEmitter() *ErrorInfoEmitter {
	return &ErrorInfoEmitter{NewSyncIndefiniteEmitter()}
}

//Implementation of IndefiniteEmitter.AddHandler.
func (e *ErrorInfoEmitter) AddHandler(handler func(interface{}, error)) {
	e.IndefiniteEmitter.AddHandler(func(args ...interface{}) {
		if args[1] == nil {
			handler(args[0], nil)
		} else {
			handler(args[0], args[1].(error))
		}
	})
}

//Implementation of IndefiniteEmitter.Emit.
func (e *ErrorInfoEmitter) Emit(i interface{}, err error) {
	e.IndefiniteEmitter.Emit(i, err)
}
