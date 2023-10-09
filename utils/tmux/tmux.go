package main

import (
	"bytes"
	"flag"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

var helpString = []string{
	"Program starts a tmux session with two windows and three panes on each window.",
	"\t Tmux shortcuts:",
	"\t\t Ctrl-b - default prefix combination",
	"\t Tmux windows:",
	"\t\t Ctrl-b c - new window",
	"\t\t Ctrl-b n, p - go to next and previus windows",
}

var paneNumbers = map[int]string{
	0: "zero",
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
}

func tmuxPrintHelp() {
	for _, msg := range helpString {
		fmt.Println(msg)
	}
}

func foo() {
}

func runSeparateOutput(cmd *exec.Cmd) (string, string, error) {
	var (
		stdout, stderr bytes.Buffer
	)

	cmd.Stderr = &stderr
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", stderr.String(), err
	}

	return stdout.String(), "", nil
}

func checkTmuxSessions() []string {
	var (
		stdout, stderr string
		sessNames      []string
	)

	tmux := exec.Command("tmux", "ls")

	stdout, stderr, err := runSeparateOutput(tmux)

	if err != nil {
		fmt.Print(stderr)
		return sessNames
	}

	sessions := strings.Split(stdout, "\n")

	for _, ses := range sessions {
		name := strings.Split(ses, ":")[0]
		sessNames = append(sessNames, name)
	}

	return sessNames
}

func startThreePane(number int) {
	var sname string

	if pNum, ok := paneNumbers[number]; ok {
		sname = "three-pane-" + pNum
	} else {
		fmt.Printf("Can't start tmux session with number - %v\n", number)
		return
	}

	tmux := exec.Command("gnome-terminal", "--full-screen", "--", "tmux", "new", "-s", sname)
	if err := tmux.Run(); err != nil {
		fmt.Println(err.Error())
		return
	}

	exec.Command("tmux", "split-window", "-h", "-t", sname).Output()
	exec.Command("tmux", "split-window", "-v", "-t", sname).Output()

	exec.Command("tmux", "new-window", "-t", sname).Output()

	exec.Command("tmux", "split-window", "-h", "-t", sname+":1").Output()
	exec.Command("tmux", "split-window", "-v", "-t", sname+":1").Output()

	exec.Command("tmux", "send-keys", "-t", sname+":1.1", "gotop", "Enter").Output()
	exec.Command("tmux", "send-keys", "-t", sname+":1.2", "mc", "Enter").Output()
}

func main() {
	pause := flag.Uint("p", 0, "Pause (in seconds) before start.")
	help := flag.Bool("h", false, "Print tmux help.")
	flag.Parse()

	if *help {
		tmuxPrintHelp()
		return
	}

	time.Sleep(time.Duration((*pause)) * time.Second)

	foo := checkTmuxSessions()
	startThreePane(len(foo))
}
