package github

import "encoding/json"

type Gist struct {
	Id          string    `json:"id"`
	Url         string    `json:"html_url"`
	Description string    `json:"description"`
	Files       gistFiles `json:"files"`
	Public      bool      `json:"public"`
}

type gistFile struct {
	Name string `json:"filename"`
	Url  string `json:"raw_url"`
}

type gistFiles []gistFile

func (gs *gistFiles) UnmarshalJSON(b []byte) error {
	m := map[string]gistFile{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	*gs = gistFiles{}
	for _, v := range m {
		*gs = append(*gs, v)
	}

	return nil
}
