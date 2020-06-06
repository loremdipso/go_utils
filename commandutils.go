package go_utils

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/kballard/go-shellquote"
)

func ExecuteCommand(command string, doWait bool) int {
	words, err := shellquote.Split(command)
	if err != nil {
		// TODO: handle this
		return -1
	}

	var cmd *exec.Cmd
	if len(words) == 1 {
		cmd = exec.Command(command)
	} else {
		cmd = exec.Command(words[0], words[1:]...)
	}

	cmd.Stdout = nil
	cmd.Stderr = nil

	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	pid := cmd.Process.Pid
	if doWait {
		cmd.Wait()
	}

	return pid
}

func ExecuteCommandAndGetResults(command string) (string, error) {
	words, err := shellquote.Split(command)
	if err != nil {
		// TODO: handle this better
		return "", err
	}

	var cmd *exec.Cmd
	if len(words) == 1 {
		cmd = exec.Command(command)
	} else {
		cmd = exec.Command(words[0], words[1:]...)
	}

	output, err := cmd.Output()
	return strings.Trim(string(output), "\n"), err
}
