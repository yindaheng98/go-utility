# github.com/yindaheng98/go-utility/Emitter

A series of event emitter, both sync and async.

## Introduction

### `interface`s in this package

There are two types of emitters defined in this package:

* A single-payload emitter can carry only 1 event payload at once
* A multi-payload emitter can carry more than 1 event payload at once

Both this two emitter will have 4 methods:

1. `Enable()`: before call this method, emitter will not run (that means the method `emit(...)` will do nothing)
2. `Disable()`: after call this method, emitter will stop running
3. `AddHandler(func(...))`: add a handler to the emitter. Handler is a function. When emitter is running and the method `Emit(...)` is called, all the handler added by `AddHandler(func(...))` will be called, and `Emit(...)`'s argument will become handler's argument
4. `Emit(...)`: receive payload and emit a event

#### `Emitter`

The interface `Emitter` is the base interface of single-payload emitter. It looks like this:

```go
type Emitter interface {
	Enable()
	Disable()
	AddHandler(func(interface{}))
	Emit(interface{})
}
```

As you see, both single-payload emitter's handler and method `Emit` can only receive 1 parameter.

#### `IndefiniteEmitter`

The interface `IndefiniteEmitter` is the base interface of multi-payload emitter. It looks like this:

```go
type IndefiniteEmitter interface {
	Enable()
	Disable()
	AddHandler(func(...interface{}))
	Emit(...interface{})
}
```

As you see, multi-payload emitter use indefinite parameter expression `...interface{}` to receive more than 1 event payload.

### `struct`s in this package

#### `SyncEmitter` and `AsyncEmitter`

Both `SyncEmitter` and `AsyncEmitter` are the basic implementation of `Emitter`. In `SyncEmitter`'s method `Emit`, handlers will run synchronously, while in `SyncEmitter`'s `Emit` handlers will run asynchronously.

PS: If you want to use thread synchronization tools such as thread lock in your handler functions, `AsyncEmitter` is recommended, because using `SyncEmitter` can easily cause synchronization problems when you put a thread lock in its handler function.

#### `SyncIndefiniteEmitter` and `AsyncIndefiniteEmitter`

Both `SyncIndefiniteEmitter` and `AsyncIndefiniteEmitter` are the basic implementation of `IndefiniteEmitter`. Just like `SyncEmitter` and `AsyncEmitter`, `SyncIndefiniteEmitter` and `AsyncIndefiniteEmitter` will run handlers synchronously and asynchronously respectively. `SyncIndefiniteEmitter` is also not recommended when you put a thread lock in its handler function.

#### `ErrorEmitter`

`ErrorEmitter` is an implementation of `Emitter`, designed for `Emit` error. It is also a simple example for how to make your own emitter using `SyncEmitter` and `AsyncEmitter`.

#### `ErrorInfoEmitter`

`ErrorInfoEmitter` is an implementation of `IndefiniteEmitter`, designed for `Emit` an error and an additional information. It is also a simple example for how to make your own emitter using `SyncIndefiniteEmitter` and `AsyncIndefiniteEmitter`.

### functions

#### `func Cascade(from Emitter, to Emitter)`

Received two emitters `from` and `to`, the function will cascade the emitters, that means, when a event was emitted in `from`, the same event will be also emit in `to`. It is very simple:

```go
func Cascade(from Emitter, to Emitter) {
	from.AddHandler(func(i interface{}) {
		to.Emit(i)
	})
}
```

#### `func IndefiniteCascade(from IndefiniteEmitter, to IndefiniteEmitter)`

Cascade for `IndefiniteEmitter`, just like `func Cascade(from Emitter, to Emitter)`, also very simple:

```go
func IndefiniteCascade(from IndefiniteEmitter, to IndefiniteEmitter) {
	from.AddHandler(func(args ...interface{}) {
		to.Emit(args...)
	})
}
```

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