package client

func (c *steamSpyClient) GetTop100Owned() ([]SteamItem, error) {
	return c.getItems(getTop100Owned, nil)
}
