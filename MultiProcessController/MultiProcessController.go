package MultiProcessController

import "sync"

//Provide the command for starting and stopping.
type Commander interface {
	//Generate the command for starting the No.i process.
	GenerateStartCommand(i uint64) (name string, args []string)
	//Generate the command for stopping the No.i process.
	GenerateStopCommand(i uint64) (name string, args []string)
}

//The major struct of MultiProcessController.
type MultiProcessController struct {
	processes []*process
	commander Commander
	mu        *sync.Mutex
	logger    Logger
}

//New returns the pointer to a new MultiProcessController struct.
func New(commander Commander, logger Logger) *MultiProcessController {
	return &MultiProcessController{[]*process{}, commander, new(sync.Mutex), logger}
}

//Start N processes.
//If there are more than N processes is running, stop those redundant processes.
func (mpc *MultiProcessController) StartN(N uint64) {
	mpc.mu.Lock()
	defer mpc.mu.Unlock()
	length := uint64(len(mpc.processes))
	if length < N { //少则补
		mpc.processes = append(mpc.processes, make([]*process, N-length)...)
		for i := length; i < N; i++ {
			nameStart, argsStart := mpc.commander.GenerateStartCommand(i)
			nameStop, argsStop := mpc.commander.GenerateStopCommand(i)
			process := newProcess(nameStart, argsStart, nameStop, argsStop, iLogger{mpc.logger, i})
			mpc.processes[i] = process
		}
	}
	//全部启动
	for i := uint64(0); i < N; i++ {
		mpc.processes[i].Start()
	}
}

//Stop all the processes.
func (mpc *MultiProcessController) StopAll() {
	mpc.mu.Lock()
	defer mpc.mu.Unlock()
	wg := new(sync.WaitGroup)
	wg.Add(len(mpc.processes))
	for _, p := range mpc.processes {
		go func(p *process) {
			p.Stop()
			wg.Done()
		}(p)
	}
	wg.Wait()
}
