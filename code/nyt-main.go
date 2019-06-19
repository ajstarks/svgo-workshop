func main() {
	var section = flag.String("h", "u.s.", usage)
	flag.Parse()
	canvas := svg.New(os.Stdout)
	canvas.Start(1200, 900)
	canvas.Rect(0, 0, 1200, 900)
	nytheadlines(canvas, *section)
	canvas.End()
}
