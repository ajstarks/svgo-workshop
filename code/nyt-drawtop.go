// drawStories walks the result structure,
// displaying title and thumbnail in a grid
func drawStories(canvas *svg.SVG, section string, width int, data NYTStories, n int) {
	top, left := 100, 150
	titley := top - 50
	x, y := left, top
	ts := fmt.Sprintf(headfmt, section, time.Now().Format(datefmt))
	if n > data.StoryCount {
		n = data.StoryCount
	}
