package client

func (c *steamSpyClient) GetTop100In2Weeks() ([]SteamItem, error) {
	return c.getItems(getTop100In2Weeks, nil)
}
