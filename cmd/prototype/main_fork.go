package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	fmt.Printf("%#v", os.Args)
	// err := syscall.Exec(os.Args[1], os.Args[2:], nil)
	proc := findProcess(os.Args[1])
	if proc != nil {
		cli, _ := proc.Cmdline()
		fmt.Println("found", proc, cli)
		return
	}

	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	cmd := exec.CommandContext(ctx, os.Args[1], os.Args[2:]...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
	// r, e := cmd.CombinedOutput()
	cmd.Start()
	fmt.Println("started")

	// fmt.Println(string(r), e, cmd.ProcessState)
	// pid, err := syscall.ForkExec(os.Args[1], os.Args[2:], &syscall.ProcAttr{
	// 	Dir:   "",
	// 	Env:   []string{},
	// 	Files: []uintptr{},
	// 	Sys:   &syscall.SysProcAttr{},
	// })
	// fmt.Println("done", pid, err)
	// time.Sleep(time.Second * 10)

}

func findProcess(name string) *process.Process {
	processes, err := process.Processes()
	if err != nil {
		return nil
	}

	for _, p := range processes {
		processName, err := p.Name()
		if err != nil {
			// return false, err
			continue
		}

		if processName == name {
			println("found", name, processName)
			return p
		}
	}

	return nil
}
