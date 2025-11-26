package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	tmux, err := exec.LookPath("tmux")
	if err != nil {
		fmt.Println(err)
	}

	cmd := exec.Command(tmux, "ls")
	if err != nil {
		fmt.Println(err)
	}

	var stdout, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	
	err = cmd.Run()

	fmt.Printf("stdout: %s\nstderr: %s\n", stdout.String(), stderr.String())
}
