// nytheadlines retrieves data from the NYTimes API, decodes and displays it.
func nytheadlines(canvas *svg.SVG, section string) {
	// Read
	r, err := netread(fmt.Sprintf(NYTfmt+param, section, NYTAPIkey))
	if err != nil {
		fmt.Fprintf(os.Stderr, "headline read error: %v\n", err)
		return
	}
	defer r.Close()
	// Decode
	var data NYTHeadlines
	if err = json.NewDecoder(r).Decode(&data); err != nil {
		fmt.Fprintf(os.Stderr, "decode: %v\n", err)
		return
	}
	// drawing code next...
}
