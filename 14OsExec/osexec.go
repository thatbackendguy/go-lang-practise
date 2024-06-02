package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	command := "fping"
	args := []string{"192.168.0.1", "-c3", "-q"}

	execute := exec.Command(command, args...)

	output, err := execute.CombinedOutput()

	if err != nil {
		fmt.Println(err)
	}

	if strings.Contains(string(output), "/0%") {
		fmt.Println("Device is UP")
	} else {
		fmt.Println("Device is DOWN")
	}
}
