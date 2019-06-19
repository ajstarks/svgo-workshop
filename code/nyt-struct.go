// NYTHeadlines is the headline info from the New York Times
type NYTHeadlines struct {
	Status     string   `json:"status"`
	Copyright  string   `json:"copyright"`
	NumResults int      `json:"num_results"`
	Results    []result `json:"results"`
}

type result struct {
	Section    string `json:"section"`
	Subsection string `json:"subsection"`
	Title      string `json:"title"`
	Abstract   string `json:"abstract"`
	Thumbnail  string `json:"thumbnail_standard"`
}
