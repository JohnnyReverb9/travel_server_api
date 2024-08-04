package test

import "travel_server_api/structs"

var Visits = map[string]structs.Visit{
	"1": {ID: 1, Location: 12, User: 10, Visited_at: 1614393810, Mark: 4},
	"2": {ID: 2, Location: 7, User: 3, Visited_at: 1613393810, Mark: 5},
	"3": {ID: 3, Location: 13, User: 10, Visited_at: 1614393813, Mark: 5},
}
