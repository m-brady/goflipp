package flipp

import (
	"net/http"
	"time"
)

// API constants
const (
	host = "https://gateflipp.flippback.com/bf/flipp"
)

type Params struct {
	PostalCode string
}

type Response struct {
	RefreshedAt string         `json:"refreshed_at,omitempty"`
	Merchants   []*Merchant    `json:"merchants,omitempty"`
	Flyers      []*Flyer       `json:"flyers,omitempty"`
	ItemDetails []*ItemDetails `json:"items,omitempty"`
	Item        *Item          `json:"item,omitempty"`
}

var flippClient = http.Client{Timeout: 30 * time.Second}
