package main

import (
	"github.com/freitzzz/pgist/internal/actions"
	"github.com/freitzzz/pgist/internal/flags"
)

func main() {
	fl := flags.ParseArgs()

	var action actions.Action
	if fl.ListGists {
		action = actions.List(fl.Username)
	}

	action.Execute()
}
