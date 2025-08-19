package client

import "errors"

func (c *steamSpyClient) GetByGenre(genre string) ([]SteamItem, error) {
	if len(genre) == 0 {
		return nil, errors.New(ValidationError_GenreIsRequired)
	}

	return c.getItems(getByGenre, map[string]string{"genre": genre})
}
