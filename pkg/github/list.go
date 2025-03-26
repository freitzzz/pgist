package github

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/freitzzz/pgist/internal/errors"
)

// Uses GitHub API to list user gists. If an API token is passed (pat), then private gists will also be
// listed (token needs to be set with proper permissions).
func ListGists(username string, pat *string) ([]Gist, error) {
	u := fmt.Sprintf("https://api.github.com/users/%s/gists", username)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, errors.ErrHttpCreateRequest
	}

	if pat != nil {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *pat))
	}

	resp, err := http.Get(u)
	if err != nil {
		return nil, errors.ErrHttpSendRequest
	}

	if resp.StatusCode == 404 {
		return nil, errors.ErrGitHubUserLookup
	}

	if resp.StatusCode != 200 {
		return nil, errors.ErrGitHubApi
	}

	gs := []Gist{}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&gs)

	if err != nil {
		return nil, errors.ErrGithHubListGistsParse
	}

	return gs, nil
}
