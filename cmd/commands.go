package cmd

import (
	"fmt"
	"strings"
	"task_traker/internal/handlers"

	"github.com/urfave/cli/v2"
)

func GetCommands() []*cli.Command {
	reader := handlers.List{}

	return []*cli.Command{
		{
			Name: "add",
			Usage: "Add a new todo to the list",
			Action: func(ctx *cli.Context) error {
				var name = ctx.Args().First()

				if (strings.TrimSpace(name) == "") {
					panic("Invalid todo name!")
				}

				result, err := reader.Add(name)
				
				if (err != nil){ 
					panic(err)
				}

				fmt.Println("Todo created successfully: ", result)

				return nil
			},
		},
	}
}