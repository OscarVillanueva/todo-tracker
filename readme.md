# Task tracker app in go

This a project described in https://roadmap.sh/projects/task-tracker

NAME:
   tracker - tracker todo list

USAGE:
   tracker [global options] command [command options]

COMMANDS:
- add               Add a new todo to the list
- delete            Remove a todo from the list
- update            Update the name of a todo
- in-progress, prg  Mark a todo as in progress
- complete          Mark a task as complete
- list              Get the list of todos
- help, h           Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help

## options of list

NAME:
   tracker list - Get the list of todos

USAGE:
   tracker list [command options]

OPTIONS:
- --completed           filter the todos by completed (default: false)
- --in-progress, --prg  filter the todos by progress (default: false)
- --todo                filter the task by todo (default: false)
- --help, -h            show help
