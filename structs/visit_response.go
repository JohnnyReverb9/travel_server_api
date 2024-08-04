package structs

type VisitResponse struct {
	Mark       uint8  `json:"mark"`
	Visited_at int64  `json:"visited_at"`
	Place      string `json:"place"`
}

type VisitsResponse struct {
	Visits []VisitResponse `json:"visits"`
}
