package flags

import (
	"flag"
	"os"

	"github.com/freitzzz/pgist/internal/errors"
)

var flagset = flag.NewFlagSet("pgist", flag.ExitOnError)
var listGists = flagset.Bool("l", false, "list user gists (requires authentication for private gists)")
var username = flagset.String("u", "", "specifies the user to operate (defaults to authenticated user)")

// Parses program arguments from package [os].
//
// Sets exit code (1) if an action can't be inferred, printing helper usage.
func ParseArgs() flags {
	err := flagset.Parse(os.Args[1:])
	errors.Swallow(err)

	fl := flags{
		ListGists: *listGists,
		Username:  *username,
	}

	if !fl.valid() {
		flagset.Usage()
		os.Exit(1)
	}

	return fl
}
