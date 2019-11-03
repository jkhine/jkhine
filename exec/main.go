package main

import (
	"fmt"
	"os/exec"
)

func main() {
	execCmd := exec.Command("ls", "-ltr")
	output, err := execCmd.CombinedOutput()
	if err != nil {
		fmt.Errorf("%s", err)
	}
	fmt.Printf("%s", output)
}