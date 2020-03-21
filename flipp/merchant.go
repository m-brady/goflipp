package flipp

import (
	"encoding/json"
	"net/http"
)

const (
	merchants = "/merchants"
)

type Merchant struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	UsBased        bool   `json:"us_based"`
	NameIdentifier string `json:"name_identifier"`
}

// Merchants list all merchants flipp has
func Merchants() (*Response, error) {
	req, err := http.NewRequest(http.MethodGet, host, nil)
	if err != nil {
		return nil, err
	}

	req.URL.Path = req.URL.Path + merchants

	resp, err := flippClient.Do(req)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(resp.Body)
	r := Response{}
	err = decoder.Decode(&r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
