package main

import (
	"os"
    "fmt"
    "paap/cmd"
)

func main() {

	switch os.Args[1] {

	case "push":
		commands := `cf login -a $CF_API -u $CF_USER -p $CF_PASS
					 cf target -o $CF_ORG -s $CF_SPACE
					 cf push -f ./bin/static-app/manifest.yml`

		cmd.RunCommands(commands)

	case "buildpacks":

		cmd.PivnetGet(
			"https://network.pivotal.io/api/v2/products/buildpacks/releases/31948/product_files/62976/download",
			"./bin/buildpacks/java-buildpack-offline-v4.8.zip")

		commands := `cf buildpacks`

		cmd.RunCommands(commands)


	case "help":
		fmt.Println("\nSETUP: You must set the following environment variables:")
		fmt.Println("     CF_API")
		fmt.Println("     CF_USER")
		fmt.Println("     CF_PASS")
		fmt.Println("     CF_ORG")
		fmt.Println("     CF_SPACE")
		fmt.Println("	 CF_NETWORK_TOKEN")
		fmt.Println("\nRUN: You can run the following commands:")
		fmt.Println("     push")
		fmt.Println("     buildpacks\n")
	}

}
