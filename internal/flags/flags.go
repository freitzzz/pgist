package flags

type flags struct {
	ListGists bool
	Username  string
}

func (f flags) valid() bool {
	if !f.ListGists {
		return false
	}

	return true
}
