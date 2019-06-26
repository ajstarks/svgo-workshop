// NYTStories is the headline info from the New York Times
type NYTStories struct {
	StoryCount int      `json:"num_results"`
	Results    []result `json:"results"`
}
type result struct {
	Title      string       `json:"title"`
	URL        string       `json:"url"`
	Multimedia []multimedia `json:"multimedia"`
}
type multimedia struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Format string `json:"format"`
}
