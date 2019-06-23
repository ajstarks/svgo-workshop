// nytStories retrieves Top Stories data
// from the New York Times API, decodes and displays it.
func nytStories(canvas *svg.SVG, section string, n int) {
	r, err := netread(fmt.Sprintf(NYTfmt, section, NYTAPIkey))
	if err != nil {
		fmt.Fprintf(os.Stderr, "headline read error: %v\n", err)
		return
	}
	defer r.Close()
	var data NYTStories
	if err = json.NewDecoder(r).Decode(&data); err != nil {
		fmt.Fprintf(os.Stderr, "decode: %v\n", err)
		return
	}
	drawStories(canvas, data, n)
}
