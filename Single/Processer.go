package Single

import "sync"

//Processor is the single task loop runner based on Thread.
//It can ensure that only one task loop is running anytime.
type Processor struct {
	thread    *Thread
	started   bool
	startedMu *sync.Mutex

	//Callback.Started() will be called when a goroutine is successfully run by method Start(process func()).
	//
	//Callback.Stopped() will be called when a goroutine run by method Start(process func()) is stopped.
	Callback *callbacks
}

//Returns the pointer to a Processor.
func NewProcessor() *Processor {
	p := &Processor{NewThread(), false, new(sync.Mutex), newCallbacks()}
	p.thread.Callback = p.Callback
	return p
}

//Run a task loop. The input function process() will be called in a loop until Stop() is called.
//Only one task loop run by this function will run at the same time.
//The task loop run by this function will be discarded when there is another task loop is not completed.
func (p *Processor) Start(process func()) {
	p.startedMu.Lock()
	defer p.startedMu.Unlock()
	p.thread.Callback = p.Callback
	if !p.started {
		p.started = true
		go p.thread.Run(func() {
			for p.started {
				process()
			}
		})
	}
}

//Stop the task loop.
func (p *Processor) Stop() {
	p.startedMu.Lock()
	defer p.startedMu.Unlock()
	p.started = false
}
