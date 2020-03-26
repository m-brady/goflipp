package flipp

import (
	"fmt"
	"testing"
)

func TestClient_Search(t *testing.T) {
	response, _ := Search(SearchParams{
		PostalCode: "M4P1V6",
		Query:      "egg",
		Merchants:  []string{"2018"},
	})
	fmt.Println(response.ItemDetails[0])
}

func TestClient_Item(t *testing.T) {
	response, _ := GetItem(ItemParams{
		ItemID: 484777629,
	})

	fmt.Println(response.Item)

}
