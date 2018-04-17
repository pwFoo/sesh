package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func launch(args []string) int {
	commands := make([]*exec.Cmd, 0, 5)
	start := 0
	for i, arg := range args {
		if arg == "|" {
			fmt.Println(args[start])
			cmd := exec.Command(args[start], args[start+1:i]...)
			commands = append(commands, cmd)
			start = i + 1
		}
		if i == len(args)-1 {
			if len(commands) == 0 {
				return launchSingleCommand(args)
			}
			fmt.Println(args[start])
			cmd := exec.Command(args[start], args[start+1:]...)
			commands = append(commands, cmd)
		}
	}
	return 1
}

func launchSingleCommand(args []string) int {
	// Spawning and executing a process
	cmd := exec.Command(args[0], args[1:]...)
	// Setting stdin, stdout, and stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = nil // making sure the command uses the current process' environment
	timestamp := time.Now().String()
	if err := cmd.Run(); err != nil {
		fmt.Printf(ERRFORMAT, err.Error())
		return 2
	}
	HISTLINE = fmt.Sprintf("%d::%s::%s", cmd.Process.Pid, timestamp, HISTLINE)
	return 1
}
