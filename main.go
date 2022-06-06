package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
    "syscall"
)

func main() {
	for _, command := range os.Args[1:] {
		fmt.Println("Running: " + command)
		splittedCommand := strings.Split(command, " ")
		cmd := exec.Command(splittedCommand[0], splittedCommand[1:]...)
		start := time.Now()
		stdout, err := cmd.Output()
		elapsed := time.Since(start)
		if err != nil {
            if msg, ok := err.(*exec.ExitError); ok {
                fmt.Printf("Time: %s stdder:\n %s \n", elapsed, err)
                os.Exit(msg.Sys().(syscall.WaitStatus).ExitStatus())
             }
			fmt.Printf("Time: %s stdder:\n %s \n", elapsed, err)
			os.Exit(1)
		}
		fmt.Printf("Time: %s stdout:\n %s \n", elapsed, stdout)
	}

}

