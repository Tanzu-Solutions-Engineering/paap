package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"net/http"
	"log"
	"encoding/json"
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
	authHeader := fmt.Sprintf("Authorization: Bearer %s", pivnetGetAccessToken())
	args := []string{"-H", authHeader, "--create-dirs", "-L", "-C", "-", "-o", output, url}
	fmt.Printf("COMMAND: curl %s",args[2:])
	command("curl", args)
}

func pivnetGetAccessToken() string {

	var accessToken (string) = "N/A"

	type RefreshToken struct {
		RefreshToken string `json:"refresh_token"`
	}

	jsonReq := RefreshToken{os.Getenv("CF_REFRESH_TOKEN")}
	jsonValue, _ := json.Marshal(jsonReq)

	request, err := http.NewRequest("POST", "https://network.pivotal.io//api/v2/authentication/access_tokens",bytes.NewBuffer(jsonValue));
	request.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Success expected: %d", res.StatusCode)
	} else {
		var body struct {
			AccessToken string `json:"access_token"`
		}
		json.NewDecoder(res.Body).Decode(&body)
		accessToken = body.AccessToken
	}
	return accessToken
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
