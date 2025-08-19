package client

import "errors"

func (c *steamSpyClient) GetByTag(tag string) ([]SteamItem, error) {
	if len(tag) == 0 {
		return nil, errors.New(ValidationError_TagIsRequired)
	}

	return c.getItems(getByTag, map[string]string{"tag": tag})
}
