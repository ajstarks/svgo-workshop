// drawStories walks the result structure, displaying title and thumbnail
func drawStories(canvas *svg.SVG, data NYTStories, n int) {
	top, left := 200, 150
	imx, titley := left-100, top-100
	x, y, th, tw := left, top, 75, 75
	ts := fmt.Sprintf("New York Times for %s", time.Now().Format(datefmt))
	imagelink := ""
	if n > data.StoryCount {
		n = data.StoryCount
	}
