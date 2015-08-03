//
// This package is a golang exercise to enable Linux kernel trace
// when executing a command.
//

package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	FTRACE_PID = "pid"
)

var (
	debugFs = ""
)

func findDebugFs() (string, error) {
	if file, err := os.Open("/proc/mounts"); err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// init is the subsystem constructor
func init() {
	if debugFs, err = findDebugFS(); err != nil {
	}
}

type KTrace struct {
	Command   string
	TraceFile string
}

func New(cmd string) *KTrace {
}

func (*KTrace) initTrace() {
}

func (*KTrace) StartCmd() {
}
