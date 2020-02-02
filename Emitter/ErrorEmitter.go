package Emitter

type ErrorEmitter struct {
	*Emitter
}

func NewErrorEmitter() *ErrorEmitter {
	return &ErrorEmitter{NewEmitter()}
}

type element struct {
	i interface{}
	e error
}

func (e *ErrorEmitter) AddHandler(handler func(interface{}, error)) {
	e.Emitter.AddHandler(func(i interface{}) {
		el := i.(element)
		handler(el.i, el.e)
	})
}

func (e *ErrorEmitter) Emit(i interface{}, err error) {
	e.Emitter.Emit(element{i, err})
}
