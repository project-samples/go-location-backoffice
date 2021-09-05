package location

import "github.com/core-go/search"

type LocationSM struct {
	*search.SearchModel
	LocationId   string `json:"locationId"`
	Type         string `json:"type"`
	LocationName string `json:"locationName`
}
