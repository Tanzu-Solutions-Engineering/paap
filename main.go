package main

import (
	"os"
    "paap/commands"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "paap"
	app.Version = "0.0.1"
	app.Description = "Demo CLI for cloud ops"

	app.Commands = append(app.Commands, commands.BuildpackCommands...)


	app.Run(os.Args)
}
