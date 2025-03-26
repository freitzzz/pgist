package actions

import (
	"fmt"

	"github.com/freitzzz/pgist/internal/errors"
	"github.com/freitzzz/pgist/internal/host"
	"github.com/freitzzz/pgist/pkg/github"
)

type listAction struct {
	Username string
}

func (c listAction) Execute() {
	var u, pat *string

	u = &c.Username
	if len(*u) == 0 {
		cs, err := host.GithubAuthenticatedUser()
		if err != nil {
			panic(err)
		}

		u = &cs.Username
		pat = &cs.PAT
	}

	gs, err := github.ListGists(*u, pat)

	if errors.Is(err, errors.ErrGitHubUserLookup) || len(gs) == 0 {
		fmt.Printf("No gist data found for user '%s'\n", c.Username)
		return
	}

	if err != nil {
		panic(err)
	}

	for _, g := range gs {
		fmt.Printf("%s: %s\n", g.Id, g.Description)
	}
}
