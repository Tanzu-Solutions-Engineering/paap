package main

import (
	"os"
    "fmt"
    "paap/cmd"
    "paap/smokehttp"
)

func main() {

	switch os.Args[1] {

	case "login" :
		cmd.RunCommands(`cf login -a $CF_API -u $CF_USER -p $CF_PASS --skip-ssl-validation`)

	case "create-lob":
		commands := `cf org demo
					cf create-quota small-org -m 10G -a 10 -r 100
					cf create-org demo -q small-org
					cf create-space development -o demo
					cf create-space production -o demo
					cf org demo
		            cf quota small-org`

		cmd.RunCommands(commands)

	case "teardown":
		cmd.RunCommands(`cf delete-org -f $CF_ORG
								   cf delete-buildpack -f java_buildpack-v48`)

		if (len(os.Args) > 2){
			cmd.RunCommands(`cf uninstall-plugin buildpack-usage
			                           cf uninstall-plugin do-all
			                           cf uninstall-plugin top`)
		}

	case "deploy-app":

		var commands string

		if (os.Getenv("CF_SPACE") == "development"){
			commands = `cf target -o $CF_ORG -s $CF_SPACE
					    cf push -f ./bin/springboot-app/manifest-development.yml`
		}

		if (os.Getenv("CF_SPACE") == "production"){
			commands = `cf target -o $CF_ORG -s $CF_SPACE
					    cf push -f ./bin/springboot-app/manifest-production.yml`
		}

		cmd.RunCommands(commands)

	case "run-smoketest" :
		smokehttp.SmokeHttp("http://springboot-app-development.local.pcfdev.io", 3)

	case "install-plugins":
		commands := `cf plugins
		             cf install-plugin -f -r CF-Community buildpack-usage
		             cf install-plugin -f -r CF-Community do-all
					 cf install-plugin -f -r CF-Community top
		             cf plugins`

		cmd.RunCommands(commands)

	case "upgrade-middleware":
		cmd.PivnetGet(
			"https://network.pivotal.io/api/v2/products/buildpacks/releases/31948/product_files/62976/download",
			"./bin/buildpacks/java-buildpack-offline-v4.8.zip")

		commands := `cf create-buildpack --enable java_buildpack-v48 ./bin/buildpacks/java-buildpack-offline-v4.8.zip 1
		             cf buildpacks
		             cf target -o $CF_ORG -s $CF_SPACE
					 cf do-all restage {}
		             cf buildpack-usage -b java_buildpack
		             cf buildpack-usage -b java_buildpack-v48`

		cmd.RunCommands(commands)

	case "help":
		fmt.Println("\nSETUP: You must set environment variables. ")
		fmt.Println("Refer to env_development for guidance")
		fmt.Println("     CF_API")
		fmt.Println("     CF_USER")
		fmt.Println("     CF_PASS")
		fmt.Println("     CF_ORG")
		fmt.Println("     CF_SPACE")
		fmt.Println("	 CF_NETWORK_TOKEN")
		fmt.Println("\nRUN: You can run the following commands:")
		fmt.Println("     login")
		fmt.Println("     install-plugins")
		fmt.Println("     create-lob")
		fmt.Println("     deploy-app")
		fmt.Println("     run-smoketest")
		fmt.Println("     upgrade-middleware\n")
		fmt.Println("     teardown")
		fmt.Println("     teardown all")
	}

}
