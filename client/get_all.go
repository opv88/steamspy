package client

import "strconv"

func (c *steamSpyClient) GetAll(page int64) ([]SteamItem, error) {
	return c.getItems(getAll, map[string]string{"page": strconv.FormatInt(page, 10)})
}
