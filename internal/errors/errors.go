package errors

import "errors"

var ErrHostHomeDirLookup = errors.New("failed to lookup home directory")
var ErrHostGitCredentialsRead = errors.New("failed to read .git-credentials")
var ErrHostNoGitHubAuthenticatedUser = errors.New("there is no authenticated user for github on this device")

var ErrHttpCreateRequest = errors.New("failed to create http request")
var ErrHttpSendRequest = errors.New("failed to send http request")

var ErrGitHubUserLookup = errors.New("failed to lookup github user")
var ErrGitHubApi = errors.New("failed to interact with github api")
var ErrGithHubListGistsParse = errors.New("failed to parse github gists json")

// Use this function to swallow errors that don't need to be taken care of.
func Swallow(err error) { /* noop */ }
