package command

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os/exec"
)

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

func ExecResultStrArray(cmdStr string) []string {
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return nil
	}
	// str, err := ioutil.ReadAll(stdout)
	var networkList = []string{}
	outputBuf := bufio.NewReader(stdout)
	for {
		output, _, err := outputBuf.ReadLine()
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Printf("Error :%s\n", err)
			}
			break
		}
		networkList = append(networkList, string(output))
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
		return nil
	}
	return networkList
}

func ExecResultStr(cmdStr string) string {
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return ""
	}
	str, err := ioutil.ReadAll(stdout)
	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
		return ""
	}
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(str)
}
