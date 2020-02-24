package Emitter

type Emitter interface {
	Enable()
	Disable()
	AddHandler(func(interface{}))
	Emit(interface{})
}
