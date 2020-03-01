package flipp

import "github.com/go-resty/resty/v2"

const (
	merchants = "/merchants"
)

type MerchantService struct {
	client resty.Client
}

func (m *MerchantService) GetAll() ([]*Merchant, error) {
	resp, err := m.client.R().SetResult([]*Merchant{}).Get(merchants)
	if err != nil {
		return nil, err
	}
	return resp.Result().([]*Merchant), nil
}

type Merchant struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	UsBased        bool   `json:"us_based"`
	NameIdentifier string `json:"name_identifier"`
}
