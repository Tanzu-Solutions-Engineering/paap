package main

import (
	"os"
    "fmt"
)

func main() {

	var cmds []string
	switch os.Args[1] {

	case "push":
		cmds = []string{
			"cf login -a $CF_API -u $CF_USER -p $CF_PASS",
			"cf target -o $CF_ORG -s $CF_SPACE",
			"cf push -f ./bin/static-app/manifest.yml"}

	case "buildpacks":
		cmds = []string{
			"echo TODO"}

	case "help":
		fmt.Printf("1.SETUP: You must set the following environment variables:\n")
		fmt.Printf("     CF_API\n")
		fmt.Printf("     CF_USER\n")
		fmt.Printf("     CF_PASS\n")
		fmt.Printf("     CF_ORG\n")
		fmt.Printf("     CF_SPACE\n")
		fmt.Printf("2.RUN: You can run the following commands:\n")
		fmt.Printf("     push\n")
		fmt.Printf("     buildpacks\n")
	}
	runCommands(cmds)
}
