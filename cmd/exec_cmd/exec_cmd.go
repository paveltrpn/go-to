package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func runAsRun(fname string) {
	cmd := exec.Command(fname)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(out.String()))
}

func runAsOutput(fname string) {
	cmd := exec.Command(fname)

	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(out))
}

func main() {
	runAsOutput("ls")
	runAsRun("ls")
}
