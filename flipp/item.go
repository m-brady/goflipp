package flipp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	search = "/items/search"
	item   = "/items/%d"
)

type SearchParams struct {
	Params
	Query     string
	Merchants []string
}

type ItemParams struct {
	Params
	ItemID int64
}

// ItemDetails encapsulates data received from flipp search
type ItemDetails struct {
	FlyerItemID      int64     `json:"flyer_item_id"`
	FlyerID          int       `json:"flyer_id"`
	ClippingImageURL string    `json:"clipping_image_url"`
	Name             string    `json:"name"`
	CurrentPrice     float64   `json:"current_price,omitempty"`
	ValidTo          time.Time `json:"valid_to"`
	ValidFrom        time.Time `json:"valid_from"`
	MerchantID       int       `json:"merchant_id"`
}

// Item encapsulates data received from flipp item endpoint
type Item struct {
	ID             int64     `json:"id"`
	FlyerID        int       `json:"flyer_id"`
	Name           string    `json:"name"`
	CurrentPrice   float64   `json:"current_price,string"`
	MerchantID     int       `json:"merchant_id"`
	FlyerValidFrom time.Time `json:"flyer_valid_from"`
	FlyerValidTo   time.Time `json:"flyer_valid_to"`
	SKU            string    `json:"sku"`
	CutoutImageURL string    `json:"cutout_image_url"`
}

// Search flipp given the search parameters using default client
func Search(params SearchParams) (*Response, error) {
	req, err := http.NewRequest(http.MethodGet, host, nil)
	if err != nil {
		return nil, err
	}

	req.URL.Path = req.URL.Path + search
	q := req.URL.Query()
	q.Add("q", params.Query)
	q.Add("locale", params.Locale)
	q.Add("postal_code", params.PostalCode)
	q.Add("merchant", strings.Join(params.Merchants, ","))
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
	fmt.Println(resp)
	return &r, nil
}

// GetItem returns a new item
func GetItem(params ItemParams) (*Response, error) {
	req, err := http.NewRequest(http.MethodGet, host, nil)
	if err != nil {
		return nil, err
	}

	req.URL.Path = req.URL.Path + fmt.Sprintf(item, params.ItemID)
	q := req.URL.Query()
	q.Add("locale", params.Locale)
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
