# github.com/yindaheng98/go-utility/Emitter

## Quick start

### How to implement an emitter

For example, if you want to implement an emitter to transport `error`s, you should at first inherit the single-payload emitter interface `Emitter`:

```go
type ErrorEmitter struct {
	Emitter
}
```

Then overload the methods `AddHandler(func(interface{}))` and `Emit(interface{})` to change the type of payload into `error`:

```go
func (e *ErrorEmitter) AddHandler(handler func(error)) {
	e.Emitter.AddHandler(func(i interface{}) {
		handler(i.(error))
	})
}

func (e *ErrorEmitter) Emit(err error) {
	e.Emitter.Emit(err)
}
```

That is how an emitter is implemented.

### How to instantiate a emitter

After implementing, a emitter should be instantiated before using. You can instantiate an emitter as a `SyncEmitter` or an `AsyncEmitter` as you wash, because both of them are implementation of `Emitter`.

If you want a `SyncEmitter` instance:

```go
emitter := ErrorEmitter{NewSyncEmitter()}
```

Or if you want a `AsyncEmitter` instance:

```go
emitter := ErrorEmitter{NewAsyncEmitter()}
```

### How to use a emitter

Now you have got a instance `emitter`, you should first add some handlers:

```go
emitter.AddHandler(func(err error) {
	fmt.Print("A handler - A error occurred:")
	fmt.Println(err)
})
emitter.AddHandler(func(err error) {
	fmt.Println("Another handler")
})
```

Then enable it before emit:

```go
emitter.Enable()
emitter.Emit(errors.New("I'm an error"))
```

Run and you will see the output.

### IndefiniteEmitter usage

`IndefiniteEmitter` is similiar to `Emitter`, the only difference is in the number of event payload. For example, if you want to transport not only `error`s, but also some additional information, you can do like this:

#### Implement

```go
type ErrorInfoEmitter struct {
	IndefiniteEmitter
}
func (e *ErrorInfoEmitter) AddHandler(handler func(interface{}, error)) {
	e.IndefiniteEmitter.AddHandler(func(args ...interface{}) {
		handler(args[0], args[1].(error))
	})
}
func (e *ErrorInfoEmitter) Emit(i interface{}, err error) {
	e.IndefiniteEmitter.Emit(i, err)
}
```

#### Instantiate

If you want a `SyncEmitter` instance:

```go
emitter := ErrorInfoEmitter{NewSyncEmitter()}
```

Or if you want an `AsyncEmitter` instance:

```go
emitter := ErrorInfoEmitter{NewAsyncEmitter()}
```

#### Use it

Add some handler:

```go
emitter.AddHandler(func(i interface{}, err error) {
	fmt.Print("A handler - A error occurred:")
	fmt.Print(err)
	fmt.Print("And here is some additional information:")
	fmt.Println(i)
})
emitter.AddHandler(func(i interface{}, err error) {
	fmt.Println("Another handler")
})
```

Run it:

```go
emitter.Enable()
emitter.Emit(errors.New("I'm an error"))
```