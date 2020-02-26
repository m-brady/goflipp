package flipp

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

// API constants
const (
	host = "https://gateflipp.flippback.com/bf/flipp"
)

type Response struct {
	RefreshedAt string     `json:"refreshed_at,omitempty"`
	Merchants   []Merchant `json:"merchants,omitempty"`
	Flyers      []*Flyer   `json:"flyers,omitempty"`
	Items       []*Item    `json:"items,omitempty"`
	Item        *Item      `json:"items,omitempty"`
}

type Config struct {
	Locale     string
	PostalCode string
}

func New(config *Config) *Client {
	return &Client{client: resty.New().SetHostURL(host).SetQueryParam("locale", config.Locale).SetQueryParam("postal_code", config.PostalCode)}
}

func NewWithClient(config *Config, hc *http.Client) *Client {
	return &Client{client: resty.NewWithClient(hc).SetHostURL(host).SetQueryParam("locale", config.Locale).SetQueryParam("postal_code", config.PostalCode)}
}

type Client struct {
	client *resty.Client
}

func (f *Client) Merchants() []Merchant {

	get, err := f.client.R().SetResult(Response{}).Get(merchants)
	if err != nil {
		panic(err)
	}
	return get.Result().(*Response).Merchants
}

func (f *Client) Flyers() ([]*Flyer, error) {

	get, err := f.client.R().SetResult(Response{}).SetQueryParams(map[string]string{"": ""}).Get(flyers)
	if err != nil {
		return nil, err
	}
	return get.Result().(*Response).Flyers, nil
}

func (f *Client) Items(flyerId int) ([]*Item, error) {

	get, err := f.client.R().SetResult(Response{}).Get(fmt.Sprintf(items, flyerId))
	if err != nil {
		return nil, err
	}
	return get.Result().(*Response).Items, nil
}

func (f *Client) Item(itemID int) (*Item, error) {

	get, err := f.client.R().SetResult(Response{}).Get(fmt.Sprintf(item, itemID))
	if err != nil {
		return nil, err
	}
	return get.Result().(*Response).Item, nil
}
