// API Info and formats
const (
	NYTAPIkey = "7f35400618208b235360ee874ed70df0"
	NYTfmt    = "http://api.nytimes.com/svc/news/v3/content/all/%s/"
	param     = ".json?api-key=%s&limit=5"
	style     = "font-size:20pt;font-family:sans-serif;fill:white"
	errfmt    = "unable to get network data for %s (%s)"
	datefmt   = "Monday Jan 2, 2006"
	usage     = "(arts, health, sports, science, technology, u.s., world)"
)

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