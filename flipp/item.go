package flipp

import "time"

const (
	item = "/items/%d"
)

type Item struct {
	ID             int64     `json:"id"`
	FlyerItemID    int64     `json:"flyer_item_id"`
	FlyerID        int       `json:"flyer_id"`
	Name           string    `json:"name"`
	CurrentPrice   float64   `json:"current_price,string,omitempty"`
	ValidTo        time.Time `json:"valid_to"`
	ValidFrom      time.Time `json:"valid_from"`
	MerchantID     int       `json:"merchant_id"`
	Price          string    `json:"price"`
	ImageURL       string    `json:"cutout_image_url"`
	FlyerValidFrom time.Time `json:"flyer_valid_from"`
	FlyerValidTo   time.Time `json:"flyer_valid_to"`
	SKU            int       `json:"sku"`
	CutoutImageURL string    `json:"cutout_image_url"`
}
