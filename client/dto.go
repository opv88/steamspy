package client

import (
	"encoding/json"
	"strconv"
)

type SteamItem struct {
	AppId           int64     `json:"appId"`
	Name            string    `json:"name"`
	Developer       string    `json:"developer"`
	Publisher       string    `json:"publisher"`
	ScoreRank       ScoreRank `json:"score_rank"`
	Positive        int64     `json:"positive"`
	Negative        int64     `json:"negative"`
	UserScore       int64     `json:"userscore"`
	Owners          string    `json:"owners"`
	AverageForever  int64     `json:"average_forever"`
	AverageTwoWeeks int64     `json:"average_2weeks"`
	MedianForever   int64     `json:"median_forever"`
	MedianTwoWeeks  int64     `json:"median_2weeks"`
	Price           string    `json:"price"`
	InitialPrice    string    `json:"initialprice"`
	Discount        string    `json:"discount"`
	CCU             int64     `json:"ccu"`
}

type SteamItemDetailed struct {
	SteamItem
	Underscore int64  `json:"underscore"`
	Languages  string `json:"languages"`
	Genre      string `json:"genre"`
	Tags       Tags   `json:"tags"`
}

type ScoreRank int64
type Tags map[string]int64

func (sr *ScoreRank) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	val, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*sr = ScoreRank(val)

	return err
}

func (t *Tags) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	var val map[string]int64
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}
	*t = val

	return nil
}
