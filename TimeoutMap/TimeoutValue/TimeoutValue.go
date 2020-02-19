package TimeoutValue

import (
	"time"
)

type TimeoutValue struct {
	element Element       //值内容
	timeout time.Duration //固定超时时间

	updateChan  chan time.Duration //传递更新信息
	stopChan    chan bool          //传递停机信息
	stoppedChan chan bool          //传递已停机信号
}

func New(element Element, timeout time.Duration) *TimeoutValue {
	v := &TimeoutValue{element, timeout,
		make(chan time.Duration, 1), make(chan bool, 1), make(chan bool, 1)}
	close(v.updateChan)
	close(v.stopChan)
	close(v.stoppedChan)
	return v
}

func (v *TimeoutValue) GetElement() Element {
	return v.element
}

func (v *TimeoutValue) GetTimeout() time.Duration {
	return v.timeout
}

//启动检查线程
func (v *TimeoutValue) Run() {
	v.updateChan = make(chan time.Duration, 1)
	v.stopChan = make(chan bool, 1)
	v.stoppedChan = make(chan bool, 1)
loop:
	for {
		select {
		case v.timeout = <-v.updateChan: //正常更新是一定时间内updateChan中传入了下一次更新的时间
		case <-time.After(v.timeout): //一定时间内updateChan中没有传入下一次更新的时间
			v.element.TimeoutedHandler()
			break loop //就停止检查线程
		case <-v.stopChan: //如果传入了停止信息
			break loop //就直接停止检查线程
		}
	}
	close(v.updateChan)
	close(v.stopChan)
	v.stoppedChan <- true
	close(v.stoppedChan)
}

func (v *TimeoutValue) Update(el Element, timeout time.Duration) {
	defer func() { recover() }()
	if el != nil {
		v.element = el
	}
	v.updateChan <- timeout
	v.element.UpdatedHandler()
}

func (v *TimeoutValue) Stop() {
	defer func() { recover() }()
	v.stopChan <- true
	<-v.stoppedChan
}
