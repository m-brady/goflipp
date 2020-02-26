package flipp

const (
	flyers = "/flyers"
	items  = "/flyers/%d"
)

type Flyer struct {
	ID            int    `json:"id"`
	MerchantID    int    `json:"merchant_id"`
	ValidTo       string `json:"valid_to"`
	ValidFrom     string `json:"valid_from"`
	AvailableTo   string `json:"available_to"`
	AvailableFrom string `json:"available_from"`
	Tweeted       bool
}
