package Emitter

type ErrorInfoEmitter struct {
	IndefiniteEmitter
}

func NewAsyncErrorInfoEmitter() *ErrorInfoEmitter {
	return &ErrorInfoEmitter{NewAsyncIndefiniteEmitter()}
}
func NewSyncErrorInfoEmitter() *ErrorInfoEmitter {
	return &ErrorInfoEmitter{NewSyncIndefiniteEmitter()}
}

func (e *ErrorInfoEmitter) AddHandler(handler func(interface{}, error)) {
	e.IndefiniteEmitter.AddHandler(func(args ...interface{}) {
		handler(args[0], args[1].(error))
	})
}

func (e *ErrorInfoEmitter) Emit(i interface{}, err error) {
	e.IndefiniteEmitter.Emit(i, err)
}
