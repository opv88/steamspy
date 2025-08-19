package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type steamSpyClient struct {
}

func NewClient() ISteamSpyClient {
	return &steamSpyClient{}
}

func (_ steamSpyClient) getItems(request string, params map[string]string) ([]SteamItem, error) {
	sb := strings.Builder{}
	sb.WriteString(baseUrl)
	sb.WriteString(request)
	if params != nil {
		for k, v := range params {
			sb.WriteString(fmt.Sprintf("&%s=%s", k, v))
		}
	}

	resp, err := http.Get(sb.String())
	if err != nil {
		msg, ok := gettingErrors[request]
		if !ok {
			msg = gettingErrors[defaultCode]
		}

		return nil, fmt.Errorf(msg, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		msg, ok := gettingRespCodeErrors[request]
		if !ok {
			msg = gettingRespCodeErrors[defaultCode]
		}

		return nil, fmt.Errorf(msg, resp.Status)
	}

	bt, _ := io.ReadAll(resp.Body)

	res := make(map[string]SteamItem, 0)
	err = json.Unmarshal(bt, &res)
	if err != nil {
		msg, ok := readingBodyErrors[request]
		if !ok {
			msg = readingBodyErrors[defaultCode]
		}

		return nil, fmt.Errorf(msg, err)
	}

	if len(res) == 0 {
		return []SteamItem{}, nil
	}

	result := make([]SteamItem, 0)
	for _, item := range res {
		result = append(result, item)
	}

	return result, nil
}
