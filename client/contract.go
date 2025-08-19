package client

type ISteamSpyClient interface {
	GetTop100In2Weeks() ([]SteamItem, error)
	GetTop100Forever() ([]SteamItem, error)
	GetTop100Owned() ([]SteamItem, error)
	GetByGenre(genre string) ([]SteamItem, error)
	GetByTag(tag string) ([]SteamItem, error)
	GetAppDetails(appId int64) (SteamItemDetailed, error)
	GetAll(page int64) ([]SteamItem, error)
}
