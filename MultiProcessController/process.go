package MultiProcessController

import (
	"github.com/yindaheng98/go-utility/Single"
	"os/exec"
	"sync"
	"time"
)

type process struct {
	main   *Single.Processor
	mainWg *sync.WaitGroup

	nameStart string
	argsStart []string
	nameStop  string
	argsStop  []string
	logger    iLogger
}

func newProcess(nameStart string, argsStart []string, nameStop string, argsStop []string, logger iLogger) (p *process) {
	p = &process{
		nameStart: nameStart,
		argsStart: argsStart,
		nameStop:  nameStop,
		argsStop:  argsStop,
		logger:    logger,

		main:   Single.NewProcessor(),
		mainWg: new(sync.WaitGroup),
	}
	p.main.Callback.Started = func() {
		p.mainWg.Add(1)
	}
	p.main.Callback.Stopped = func() {
		p.mainWg.Done()
	}
	return
}

//Start the process, with a command and some args
func (p *process) Start() {
	p.main.Start(func() {
		run(exec.Command(p.nameStart, p.argsStart...), p.logger.Log) //循环运行这个指令
		time.Sleep(1e8)
	})
}

func (p *process) Stop() {
	p.Start()
	p.main.Stop()
	stopper := Single.NewProcessor()
	stopper.Start(func() {
		run(exec.Command(p.nameStop, p.argsStop...), func(i ...interface{}) {
			p.logger.Log(append([]interface{}{"stopper-->"}, i...)...)
		})
		time.Sleep(1e8)
	})
	p.mainWg.Wait()
	stopper.Stop()
}
