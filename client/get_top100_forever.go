package client

func (c *steamSpyClient) GetTop100Forever() ([]SteamItem, error) {
	return c.getItems(getTop100Forever, nil)
}
