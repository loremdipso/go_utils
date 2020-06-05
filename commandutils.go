package go_utils

import (
	"fmt"
	"os/exec"

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
