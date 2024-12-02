package client

import (
	"fmt"
	"os"
	"syscall"

	"github.com/rs/zerolog/log"
)

var (
	errFileNotFound = fmt.Errorf("file not found")
)

func dettachFork(args []string) (*os.Process, error) {
	log.Printf("dettachServer: args: %v", args)
	defer log.Printf("dettachServer end args: %v", args)
	if len(args) < 1 {
		return nil, errFileNotFound
	}
	attr := &os.ProcAttr{
		Dir:   "",
		Env:   []string{},
		Files: []*os.File{},
		Sys: &syscall.SysProcAttr{
			Setsid: true,
		},
	}

	bin := args[0]
	params := args[1:]
	return os.StartProcess(bin, params, attr)
}
