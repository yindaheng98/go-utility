/*
Package main just written for test.
*/
package main

import (
	"fmt"
	"github.com/yindaheng98/go-utility/MultiProcessController"
	"log"
	"time"
)

type TestCommander struct {
	nameStart string
	argsStart func(i uint64) []string
	nameStop  string
	argsStop  func(i uint64) []string
}

func (t TestCommander) GenerateStartCommand(i uint64) (name string, args []string) {
	name = t.nameStart
	args = t.argsStart(i)
	return
}
func (t TestCommander) GenerateStopCommand(i uint64) (name string, args []string) {
	name = t.nameStop
	args = t.argsStop(i)
	return
}

type TestLogger struct {
}

func (TestLogger) Log(args ...interface{}) {
	log.Println(args...)
}

var jmeterCommander = TestCommander{
	nameStart: "jmeter",
	argsStart: func(i uint64) []string {
		return []string{
			fmt.Sprintf("-Jserver.rmi.localport=%d", 1099+i),
			fmt.Sprintf("-Dserver_port=%d", 1099+i),
			"--server",
			"-Jserver.rmi.ssl.disable=true"}
	},
	nameStop: "sh",
	argsStop: func(i uint64) []string {
		return []string{"-c",
			fmt.Sprintf("kill $(echo `ps -ef | grep %d | awk '{print $1}'`)", 1099+i),
		}
	},
}

func main() {
	mpc := MultiProcessController.New(jmeterCommander, TestLogger{})
	mpc.StartN(3)
	time.Sleep(15e9)
	mpc.StartN(6)
	time.Sleep(15e9)
	mpc.StartN(6)
	time.Sleep(15e9)
	mpc.StopAll()
	time.Sleep(15e9)
	mpc.StartN(3)
	time.Sleep(15e9)
	mpc.StopAll()
	time.Sleep(15e9)
}
