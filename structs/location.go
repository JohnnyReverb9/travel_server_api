package structs

type Location struct {
	ID       uint32 `json:"id"`
	Place    string `json:"place"`   // text
	Country  string `json:"country"` // until 50 symbols
	City     string `json:"city"`    // until 50 symbols
	Distance uint32 `json:"distance"`
}
