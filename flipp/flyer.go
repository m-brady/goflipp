package flipp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	flyers = "/flyers/%d"
)

type Flyer struct {
	ID            int    `json:"id"`
	MerchantID    int    `json:"merchant_id"`
	ValidTo       string `json:"valid_to"`
	ValidFrom     string `json:"valid_from"`
	AvailableTo   string `json:"available_to"`
	AvailableFrom string `json:"available_from"`
}

type FlyerParams struct {
	Params
	FlyerID int
}

// Search flipp given the search parameters using default client
func Flyers(params FlyerParams) (*Response, error) {
	req, err := http.NewRequest(http.MethodGet, host, nil)
	if err != nil {
		return nil, err
	}

	req.URL.Path = req.URL.Path + fmt.Sprintf(flyers, params.FlyerID)

	q := req.URL.Query()
	q.Add("postal_code", params.PostalCode)
	req.URL.RawQuery = q.Encode()

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
