package command

import "os/exec"

func OnlyExec(cmdStr string) {
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		return
	}
	if err := cmd.Wait(); err != nil {
		return
	}
}
