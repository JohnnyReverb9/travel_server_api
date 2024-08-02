package test

import "travel_server_api/structs"

var Users = map[string]structs.User{
	"10": {ID: 10, Email: "johndoe@gmail.com", FirstName: "John", LastName: "Doe", Gender: "m", BirthDate: -1613433600},
	"3":  {ID: 3, Email: "johnny1234@gmail.com", FirstName: "Johnny", LastName: "Reverb", Gender: "m", BirthDate: -1613499600},
}
