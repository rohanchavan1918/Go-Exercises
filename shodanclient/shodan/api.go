package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIInfo struct {
	// {"member": false, "credits": 0, "display_name": null, "created": "2020-04-21T12:34:37.749000"}
	Member      bool   `json:"member"`
	Credits     int    `json:"credits"`
	DisplayName string `json:"display_name"`
	Created     string `json:"created"`
}

func (s *Client) APIInfo() (*APIInfo, error) {
	res, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseURL, s.apiKey))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret APIInfo
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
