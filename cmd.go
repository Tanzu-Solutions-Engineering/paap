package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func runCommands(cmds []string){

	for _, element := range cmds {

		var stdoutBuf, stderrBuf bytes.Buffer
		arg0 := strings.Split(element, " ")[0]
		argN := strings.Split(element," ")[1:]

		for index, anArg := range argN {
			if strings.Contains(anArg, "$") {
				argN[index] = os.Getenv(anArg[1:])
			}
		}

		fmt.Printf("COMMAND: %s\n\n",element)

		cli_cmd := exec.Command(arg0, argN...)

		stdoutIn, _ := cli_cmd.StdoutPipe()
		stderrIn, _ := cli_cmd.StderrPipe()

		var errStdout, errStderr error
		stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
		stderr := io.MultiWriter(os.Stderr, &stderrBuf)

		err := cli_cmd.Start()
		if err != nil {
			fmt.Printf("%s\n", err)
		}

		go func() {
			_, errStdout = io.Copy(stdout, stdoutIn)
		}()

		go func() {
			_, errStderr = io.Copy(stderr, stderrIn)
		}()

		err = cli_cmd.Wait()
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		if errStdout != nil || errStderr != nil {
			fmt.Printf("failed to capture stdout or stderr\n")
		}
		fmt.Printf("\n----------------------------------------------------------------------\n")
	}
}
