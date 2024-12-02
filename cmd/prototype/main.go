package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	os.Setenv("XX", "on")
	fmt.Println(os.Args)
	// return
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	os.Stdout.Write(out)
	// println(string(out))
}
