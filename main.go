package main

import (
	"os"
    "fmt"
)

func main() {

	switch os.Args[1] {

	case "push":
		commands := `cf login -a $CF_API -u $CF_USER -p $CF_PASS
					 cf target -o $CF_ORG -s $CF_SPACE
					 cf push -f ./bin/static-app/manifest.yml`

		RunCommandString(commands)

	case "buildpacks":
		commands := `echo TODO
					 echo TODO2
					 echo TODO3`

		RunCommandString(commands)

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

}
