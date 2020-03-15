package flipp

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	item   = "/items/%d"
	search = "/items/search"
)

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

func (f *Client) Search(query string, merchants ...int) ([]*ItemDetails, error) {
	var ms []string
	for _, m := range merchants {
		ms = append(ms, strconv.Itoa(m))
	}

	response, err := f.client.R().SetResult(Response{}).SetQueryParam("q", query).SetQueryParam("merchant", strings.Join(ms, ",")).Get(search)
	if err != nil {
		return nil, err
	}

	return response.Result().(*Response).ItemDetails, nil
}

func (f *Client) Item(itemID int64) (*Item, error) {

	get, err := f.client.R().SetResult(Response{}).Get(fmt.Sprintf(item, itemID))
	if err != nil {
		return nil, err
	}
	return get.Result().(*Response).Item, nil
}
