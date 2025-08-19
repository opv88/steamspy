package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *steamSpyClient) GetAppDetails(appId int64) (SteamItemDetailed, error) {
	if appId == 0 {
		return SteamItemDetailed{}, errors.New(ValidationError_AppIdIsRequired)
	}

	return c.getItemDetailed(getAppDetails, appId)
}

func (_ steamSpyClient) getItemDetailed(request string, appId int64) (SteamItemDetailed, error) {
	res := SteamItemDetailed{}

	sb := strings.Builder{}
	sb.WriteString(baseUrl)
	sb.WriteString(request)
	sb.WriteString(fmt.Sprintf("&appId=%d", appId))

	resp, err := http.Get(sb.String())
	if err != nil {
		return res, fmt.Errorf(gettingErrors[request], err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return res, fmt.Errorf(gettingRespCodeErrors[request], resp.Status)
	}

	bt, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(bt, &res)
	if err != nil {
		return res, fmt.Errorf(readingBodyErrors[request], err)
	}

	return res, nil
}
