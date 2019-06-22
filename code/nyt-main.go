func main() {
	var section = flag.String("s", "home", usage)
	var nstories = flag.Int("n", 5, "number of stories")
	flag.Parse()
	width, height := 1200, *nstories*150
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height)
	nytStories(canvas, *section, *nstories)
	canvas.End()
}
