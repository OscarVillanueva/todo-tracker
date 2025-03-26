package main

import (
	"log"
	"os"
	"task_traker/cmd"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App {
		Name: "tracker",
		Usage: "tracker todo list",
		Commands: cmd.GetCommands(),
	}

	if err := app.Run(os.Args); err != nil { 
		log.Fatal(err)
	}
	
}