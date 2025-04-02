package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"task_traker/internal/handlers"
	"task_traker/internal/models"

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
		{
			Name: "delete",
			Usage: "Remove a todo from the list",
			Action: func(ctx *cli.Context) error {
				id := ctx.Args().First()

				parsed, err := strconv.ParseInt(id, 10, 16)

				if (err != nil){
					panic("Invalid todo id!")
				}

				_, delErr := reader.Delete(int(parsed))

				if (delErr != nil){
					panic(delErr)
				}

				fmt.Println("Todo deleted successfully: ", id)

				return nil
			},
		},
		{
			Name: "update",
			Usage: "Update the name of a todo",
			Action: func(ctx *cli.Context) error {
				if (ctx.Args().Len() < 2) {
					panic("Missing arguments")
				}

				id := ctx.Args().First()
				name := ctx.Args().Get(1)

				parsed, err := strconv.ParseInt(id, 10, 16)

				if (err != nil){
					panic("Invalid todo id!")
				}

				if (strings.TrimSpace(name) == "") {
					panic("Invalid name!")
				}

				_, updateError := reader.Update(int16(parsed), name)

				if (updateError != nil){
					panic(updateError)
				}

				fmt.Printf("The task %d was updated successfully", parsed)
				return nil
			},
		},
		{
			Name: "in-progress",
			Usage: "Mark a todo as in progress",
			Aliases: []string{"prg"},
			Action: func(ctx *cli.Context) error {
				id := ctx.Args().First()

				parsed, err := strconv.ParseInt(id, 10, 16)

				if (err != nil){
					panic("Invalid todo id!")
				}

				_, prgError := reader.UpdateStatus(int16(parsed), models.IN_PROGRESS)

				if (prgError != nil) {
					panic(prgError)
				}

				fmt.Println("Mark as in progress: ", id)
				return nil
			},
		},
		{
			Name: "complete",
			Usage: "Mark a task as complete",
			Action: func(ctx *cli.Context) error {
				id := ctx.Args().First()

				parsed, err := strconv.ParseInt(id, 10, 16)

				if (err != nil){
					panic("Invalid todo id!")
				}

				_, prgError := reader.UpdateStatus(int16(parsed), models.DONE)

				if (prgError != nil) {
					panic(prgError)
				}

				fmt.Println("Mark as completed: ", id)
				return nil
			},
		},
		{
			Name: "list",
			Usage: "Get the list of todos",
			Flags: []cli.Flag {
				&cli.BoolFlag{
					Name: "completed",
					Usage: "filter the todos by completed",
				},
				&cli.BoolFlag{
					Name: "in-progress",
					Aliases: []string{"prg"},
					Usage: "filter the todos by progress",
				},
				&cli.BoolFlag{
					Name: "todo",
					Usage: "filter the task by todo",
				},
			},
			Action: func(ctx *cli.Context) error {
				var status models.Status		

    		if ctx.Bool("completed") {
					status = models.DONE
    		}

				if ctx.Bool("in-progress") {
					status = models.IN_PROGRESS	
				}

				if ctx.Bool("todo") {
					status = models.CREATED
				}

				list, err := reader.GetList(status)

				if (err != nil){
					panic(err)
				}

				bytes, bErr := json.Marshal(list)

				if (bErr != nil) {
					panic(bErr)
				}

				fmt.Printf("%s", bytes)

				return nil	
			},
		},
	}
}