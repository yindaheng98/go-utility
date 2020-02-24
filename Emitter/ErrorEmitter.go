package Emitter

type ErrorEmitter struct {
	Emitter Emitter
}

func NewAsyncErrorEmitter() *ErrorEmitter {
	return &ErrorEmitter{NewAsyncEmitter()}
}
func NewSyncErrorEmitter() *ErrorEmitter {
	return &ErrorEmitter{NewSyncEmitter()}
}

func (e *ErrorEmitter) AddHandler(handler func(error)) {
	e.Emitter.AddHandler(func(i interface{}) {
		handler(i.(error))
	})
}

func (e *ErrorEmitter) Emit(err error) {
	e.Emitter.Emit(err)
}
