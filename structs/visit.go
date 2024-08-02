package structs

type Visit struct {
	ID       uint32 `json:"id"`
	Location uint32 `json:"location"` // location id
	User     uint32 `json:"city"`     // user id
	Visited  int    `json:"visited"`  // timestamp
	Mark     uint8  `json:"mark"`     // int from 0 to 5
}
