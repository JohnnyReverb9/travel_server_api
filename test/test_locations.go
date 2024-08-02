package test

import "travel_server_api/structs"

var Locations = map[string]structs.Location{
	"12": {ID: 12, Place: "Coliseum", Country: "Italy", City: "Rome", Distance: 1024},
	"7":  {ID: 7, Place: "Pizzeria", Country: "Italy", City: "Sicilian", Distance: 14},
}
