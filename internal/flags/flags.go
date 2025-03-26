package flags

type flags struct {
	Username  string
	ListGists bool
}

func (f flags) valid() bool {
	if !f.ListGists {
		return false
	}

	return true
}
