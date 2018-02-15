package commands

import (
	"gopkg.in/urfave/cli.v1"
	"os"
	"paap/smokehttp"
	"paap/cmd"
)

var BuildpackCommands = []cli.Command{
	{
		Name:  "login",
		Usage: "login to cloud foundry using env vars for api, username, password",
		Category: "Buildpack Demo",
		Action: func(c *cli.Context) error {
			cmd.RunCommands(`cf login -a $CF_API -u $CF_USER -p $CF_PASS --skip-ssl-validation`)
			return nil
		},
	},
	{
		Name:  "create-lob",
		Usage: "create pre-defined org with 'development' and 'production' space, 10GB RAM, 10 AI quota",
		Category: "Buildpack Demo",
		Action: func(c *cli.Context) error {
			commands := `cf org demo
					cf create-quota small-org -m 10G -a 10 -r 100
					cf create-org demo -q small-org
					cf create-space development -o demo
					cf create-space production -o demo
					cf org demo
		            cf quota small-org`

			cmd.RunCommands(commands)
			return nil
		},
	},
	{
		Name:  "deploy-app",
		Usage: "deploy application to the 'development' or 'production' space based on $CF_SPACE env variable",
		Category: "Buildpack Demo",
		Action: func(c *cli.Context) error {
			var commands string
			if (os.Getenv("CF_SPACE") == "development") {
				commands = `cf target -o $CF_ORG -s $CF_SPACE
								cf push -f ./bin/springboot-app/manifest-development.yml`
			}
			if (os.Getenv("CF_SPACE") == "production") {
				commands = `cf target -o $CF_ORG -s $CF_SPACE
								cf push -f ./bin/springboot-app/manifest-production.yml`
			}

			cmd.RunCommands(commands)
			return nil
		},
	},
	{
		Name:  "run-smoketest",
		Usage: "HTTP GET smoke test that runs until error and loops every 3 seconds",
		Category: "Buildpack Demo",
		Action: func(c *cli.Context) error {
			smokehttp.SmokeHttp("http://springboot-app-development.local.pcfdev.io", 3)
			return nil
		},
	},
	{
		Name:  "install-plugins",
		Usage: "Install 'buildpack-usage', 'do-all' and 'top' cf plugins",
		Category: "Buildpack Demo",
		Action: func(c *cli.Context) error {
			commands := `cf plugins
		             cf install-plugin -f -r CF-Community buildpack-usage
		             cf install-plugin -f -r CF-Community do-all
					 cf install-plugin -f -r CF-Community top
		             cf plugins`

			cmd.RunCommands(commands)
			return nil
		},
	},
	{
		Name:  "upgrade-middleware",
		Usage: "upgrade to java v4.8 buildpack and restage all apps in $CF_SPACE",
		Category: "Buildpack Demo",
		Action: func(c *cli.Context) error {
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

			//cmd.PivnetGetToken()
			return nil
		},
	},
	{
		Name:  "teardown",
		Usage: "delete pre-defined org including spaces and applications",
		Category: "Buildpack Demo",
		Action: func(c *cli.Context) error {
			cmd.RunCommands(`cf delete-org -f $CF_ORG
									   cf delete-buildpack -f java_buildpack-v48`)
			return nil
		},
	},
}
