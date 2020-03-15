package flipp

import (
	"github.com/go-resty/resty/v2"
	"net/http"
)

// API constants
const (
	host = "https://gateflipp.flippback.com/bf/flipp"
)

type Response struct {
	RefreshedAt string         `json:"refreshed_at,omitempty"`
	Merchants   []*Merchant    `json:"merchants,omitempty"`
	Flyers      []*Flyer       `json:"flyers,omitempty"`
	ItemDetails []*ItemDetails `json:"items,omitempty"`
	Item        *Item          `json:"item,omitempty"`
}

type Config struct {
	Locale     string
	PostalCode string
}

func New(config *Config) *Client {
	return &Client{client: resty.New().SetHostURL(host).
		SetQueryParam("locale", config.Locale).
		SetQueryParam("postal_code", config.PostalCode)}
}

func NewWithClient(config *Config, hc *http.Client) *Client {
	return &Client{client: resty.NewWithClient(hc).SetHostURL(host).
		SetQueryParam("locale", config.Locale).
		SetQueryParam("postal_code", config.PostalCode)}
}

type Client struct {
	client *resty.Client
}
