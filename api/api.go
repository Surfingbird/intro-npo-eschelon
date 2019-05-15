package api

type Request struct {
	SearchText string   `json:"search-text"`
	Sites      []string `json:"sites"`
}

type Response struct {
	FoundAtSites []string `json:"found-at-sites"`
}
