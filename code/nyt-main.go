func main() {
	var section = flag.String("s", "home", usage)
	var nstories = flag.Int("n", 25, "number of stories")
	flag.Parse()

	width, height := 1000, (*nstories/5)*150
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height)
	nytStories(canvas, width, *section, *nstories)
	canvas.End()
}
