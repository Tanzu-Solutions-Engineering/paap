package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func command(cmd string, args []string){
	cli_cmd := exec.Command(cmd, args...)

	stdoutIn, _ := cli_cmd.StdoutPipe()
	stderrIn, _ := cli_cmd.StderrPipe()

	var stdoutBuf, stderrBuf bytes.Buffer
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

func PivnetGet(url, output string){
	authHeader := fmt.Sprintf("Authorization: Token %s", os.Getenv("CF_NETWORK_TOKEN"))
	args := []string{"-nc", "--header", authHeader, url, "-O", output}
	fmt.Printf("COMMAND: wget %s %s \n\n",url, output)
	command("wget", args)
}

func RunCommands(cmdString string){

	commands := strings.Split(cmdString, "\n")

	for _, element := range commands {

		element = strings.Trim(element, "\t")
		element = strings.TrimLeft(element, " ")

		arg0 := strings.Split(element, " ")[0]
		argN := strings.Split(element," ")[1:]

		for index, anArg := range argN {
			if strings.Contains(anArg, "$") {
				argN[index] = os.Getenv(anArg[1:])
			}
		}

		fmt.Printf("COMMAND: %s\n\n",element)

		command(arg0, argN)
	}
}
