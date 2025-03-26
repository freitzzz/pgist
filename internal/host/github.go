package host

import (
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/freitzzz/pgist/internal/errors"
)

type credentials struct {
	Username string
	PAT      string
}

// Tries to lookup device stored credentials (username, token) to interact with GitHub API, from the local .git-credentials file.
//
// Only entries for public GitHub are looked for (@github.com).
func GithubAuthenticatedUser() (*credentials, error) {
	d, err := os.UserHomeDir()
	if err != nil {
		return nil, errors.ErrHostHomeDirLookup
	}

	p := path.Join(d, ".git-credentials")
	bs, err := os.ReadFile(p)

	if os.IsNotExist(err) {
		return nil, errors.ErrHostNoGitHubAuthenticatedUser
	}

	if err != nil {
		return nil, errors.ErrHostGitCredentialsRead
	}

	ln := strings.Split(string(bs), "\n")
	cs := parseGitCredentials(ln)

	return cs, nil
}

func parseGitCredentials(lines []string) *credentials {
	for _, ln := range lines {
		if strings.HasPrefix(ln, "#") {
			continue
		}

		if !strings.HasSuffix(ln, "@github.com") {
			continue
		}

		u, err := url.Parse(ln)
		if err != nil {
			continue
		}

		p, ok := u.User.Password()
		if !ok {
			continue
		}

		return &credentials{
			Username: u.User.Username(),
			PAT:      p,
		}
	}

	return nil
}
