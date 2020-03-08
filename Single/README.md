# github.com/yindaheng98/go-utility/Single

This package contains following:

1. A single thread runner designed for guarante that only one thread is running at a time.
2. A single task runner based on the single thread runner, which designed for cycling through a task.

## Introduction

There are two struct `Thread` and `Processor` in this package. `Thread` is the single thread runner, and `Processor` is the single task runner.

### `Thread`

This struct has just one exported member:

```go
type Thread struct {
	Callback  *callbacks
}

type callbacks struct {
	Started func()
	Stopped func()
}
```

And just one method:

```go
func (s *Thread) Run(routine func())
```

The method can receive a function `routine`. It at first check if there is another routine is running, if so, it will directly return; if not it will at first call the start callback `s.Callback.Started()`, then run the `routine`, and call the `s.Callback.Stopped()` when `routine` returned.

If you call the `Run` of a `Thread` in different goroutine at the same time, only one will run the `routine`, all the others will return directly.

### `Processor`

This struct is based on `Thread`. Its exported member is exactly same as `Thread`. Difference is that `Processor` has two methods:

```go
func (p *Processor) Start(process func())
```

```go
func (p *Processor) Stop()
```

The method `Start(process func())` received a function `process`. It will make a goroutine cycling through the `process()`, and run the goroutine using an unexported `Thread` member, and the method `Stop()` is used for stop this goroutine. Just like the `Thread.Run(routine func())`, `p.Start(process func())` will callback `p.Callback.Started()` at the beginning of the goroutine, and call the `p.Callback.Stopped()` when stopped.

## Quick start

`Single_test.go` is two simple example of `Thread` and `Processor`, have fun hacking them.