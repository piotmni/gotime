package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	for _, command := range os.Args[1:] {
		fmt.Println("Running: " + command)
		splittedCommand := strings.Split(command, " ")
		cmd := exec.Command(splittedCommand[0], splittedCommand[1:]...)
		start := time.Now()
		stdout, stderr := cmd.Output()
		elapsed := time.Since(start)
		if stderr != nil {
			fmt.Printf("Time: %s stdder:\n %s \n", elapsed, stderr)
			os.Exit(1)
		}
		fmt.Printf("Time: %s stdout %s \n", elapsed, stdout)
	}

}

