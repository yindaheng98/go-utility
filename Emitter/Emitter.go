package Emitter

//This is the basic interface of single-payload emitter.
//This type of emitter can carry only 1 payload at a event.
type Emitter interface {

	//Enable the emitter.
	Enable()

	//Disable the emitter.
	Disable()

	//Add a handler to the emitter.
	//
	//Handler is a function.
	//When `Emit(...)` is called, all the handler added by `AddHandler(func(...))` will be called,
	//and `Emit(...)`'s argument will become handler's argument.
	AddHandler(func(interface{}))

	//Receive payload and emit a event.
	Emit(interface{})
}
