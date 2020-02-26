package Emitter

func Cascade(from Emitter, to Emitter) {
	from.AddHandler(func(i interface{}) {
		to.Emit(i)
	})
}

func IndefiniteCascade(from IndefiniteEmitter, to IndefiniteEmitter) {
	from.AddHandler(func(args ...interface{}) {
		to.Emit(args...)
	})
}
