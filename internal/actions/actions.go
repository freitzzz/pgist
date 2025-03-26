package actions

type Action interface {
	Execute()
}

// Creates an action for listing gists.
func List(username string) listAction {
	return listAction{
		Username: username,
	}
}
