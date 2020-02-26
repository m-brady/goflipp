package flipp

const (
	merchants = "/merchants"
)

type Merchant struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	UsBased        bool   `json:"us_based"`
	NameIdentifier string `json:"name_identifier"`
}
