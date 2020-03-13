package Emitter

//Cascade implementations of `Emitter` "from" and "to".
//After cascaded, when a event was emitted in `from`,
//the same event will be emit in `to`.
func Cascade(from Emitter, to Emitter) {
	from.AddHandler(func(i interface{}) {
		to.Emit(i)
	})
}

//Cascade implementations of `IndefiniteEmitter` "from" and "to".
//After cascaded, when a event was emitted in `from`,
//the same event will be emit in `to`.
func IndefiniteCascade(from IndefiniteEmitter, to IndefiniteEmitter) {
	from.AddHandler(func(args ...interface{}) {
		to.Emit(args...)
	})
}
