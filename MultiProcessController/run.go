package MultiProcessController

import (
	"bufio"
	"context"
	"os/exec"
	"regexp"
	"time"
)

var rex, _ = regexp.Compile("\\s")

//运行cmd直到它退出
func run(cmd *exec.Cmd, log func(...interface{})) {
	log("Running: ", cmd.Args)

	stdout, err := cmd.StdoutPipe()
	if err != nil { //获取输出
		log(err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background()) //用于标识进程退出
	defer cancel()
	buf := bufio.NewReader(stdout)
	go func() { //命令行输出线程
		for {
			line, _, _ := buf.ReadLine()
			s := string(line)
			if len(rex.ReplaceAllString(s, "")) > 0 {
				log(s)
			}
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(1e8)
				continue
			}
		}
	}()

	if err := cmd.Start(); err != nil {
		log(err)
	}
	if err := cmd.Wait(); err != nil {
		log(err)
	}
}
