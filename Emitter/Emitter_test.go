package Emitter

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type event struct {
	name string
}

func TestEmitter(t *testing.T) {
	emitter := NewAsyncEmitter()
	emitter.AddHandler(func(e interface{}) {
		t.Log("Here is a handler, I'm handling: " + e.(event).name)
	})
	emitter.Enable()
	go emitter.Emit(event{"I'm a event"})
	emitter.AddHandler(func(e interface{}) {
		t.Log("Here is another handler, I'm handling: " + e.(event).name)
	})
	go emitter.Emit(event{"I'm another event"})
	go emitter.Enable()
	go emitter.Emit(event{"I'm event2"})
	emitter.AddHandler(func(e interface{}) {
		t.Log("Here is handler2, I'm handling: " + e.(event).name)
	})
	go emitter.Emit(event{"I'm event3"})
	go emitter.Disable()
	go emitter.Enable()
	go emitter.Emit(event{"I'm event4"})
	go emitter.Disable()
	emitter.AddHandler(func(e interface{}) {
		t.Log("Here is handler3, I'm handling: " + e.(event).name)
	})
	go emitter.Emit(event{"I'm event5"})
	go emitter.Disable()
	time.Sleep(1e9 * 3)
}

func TestErrorInfoEmitter(t *testing.T) {
	emitter := NewAsyncErrorInfoEmitter()
	emitter.AddHandler(func(e interface{}, err error) {
		t.Log(fmt.Sprintf("Here is a handler, I'm handling: %s, and the error is %s", e.(event).name, err.Error()))
	})
	emitter.Enable()
	go emitter.Emit(event{"I'm a event"}, errors.New("i'm an error"))
	emitter.AddHandler(func(e interface{}, err error) {
		t.Log(fmt.Sprintf("Here is another handler, I'm handling: %s, and the error is %s", e.(event).name, err.Error()))
	})
	go emitter.Emit(event{"I'm another event"}, errors.New("i'm another error"))
	go emitter.Enable()
	go emitter.Emit(event{"I'm event2"}, errors.New("i'm error2"))
	emitter.AddHandler(func(e interface{}, err error) {
		t.Log(fmt.Sprintf("Here is handler2, I'm handling: %s, and the error is %s", e.(event).name, err.Error()))
	})
	go emitter.Emit(event{"I'm event3"}, errors.New("i'm error3"))
	go emitter.Disable()
	go emitter.Enable()
	go emitter.Emit(event{"I'm event4"}, errors.New("i'm error4"))
	go emitter.Disable()
	emitter.AddHandler(func(e interface{}, err error) {
		t.Log("Here is handler3, I'm handling: " + e.(event).name)
	})
	go emitter.Emit(event{"I'm event5"}, errors.New("i'm error5"))
	go emitter.Disable()
	time.Sleep(1e9 * 3)
}
