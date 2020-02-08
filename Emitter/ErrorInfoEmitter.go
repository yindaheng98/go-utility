package Emitter

type ErrorInfoEmitter struct {
	*Emitter
}

func NewErrorInfoEmitter() *ErrorInfoEmitter {
	return &ErrorInfoEmitter{NewEmitter()}
}

type element struct {
	i interface{}
	e error
}

func (e *ErrorInfoEmitter) AddHandler(handler func(interface{}, error)) {
	e.Emitter.AddHandler(func(i interface{}) {
		el := i.(element)
		handler(el.i, el.e)
	})
}

func (e *ErrorInfoEmitter) Emit(i interface{}, err error) {
	e.Emitter.Emit(element{i, err})
}
