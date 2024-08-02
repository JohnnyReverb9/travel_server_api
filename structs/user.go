package structs

type User struct {
	ID        uint32 `json:"id"`
	Email     string `json:"email"`      // until 100 symbols
	FirstName string `json:"first_name"` // until 50 symbols
	LastName  string `json:"last_name"`  // until 50 symbols
	Gender    string `json:"gender"`     // m -- male; f -- female
	BirthDate int    `json:"birth_date"` // timestamp
}
