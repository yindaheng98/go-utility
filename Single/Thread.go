package Single

import "sync"

//Thread is the basic of package Single.
//It can ensure that only one goroutine is running anytime.
type Thread struct {
	started   bool
	startedMu *sync.Mutex

	//Callback.Started() will be called when a goroutine is successfully run by method Run(routine func()).
	//
	//Callback.Stopped() will be called when a goroutine run by method Run(routine func()) is stopped.
	Callback *callbacks
}

//Returns the pointer to a Thread.
func NewThread() *Thread {
	return &Thread{false, new(sync.Mutex), newCallbacks()}
}

//Run a goroutine.
//Only one goroutine run by this function will run at the same time.
//The goroutine run by this function will be discarded when there is another goroutine is not completed.
func (s *Thread) Run(routine func()) {
	s.startedMu.Lock()
	if s.started { //如果已经启动过了
		s.startedMu.Unlock()
		return //就直接返回
	}
	s.started = true //否则就进入已启动状态
	s.Callback.Started()
	s.startedMu.Unlock()
	defer func() {
		s.startedMu.Lock()
		s.started = false //在程序退出时重新回到未启动状态
		s.Callback.Stopped()
		s.startedMu.Unlock()
	}()
	routine() //然后启动协程
}
