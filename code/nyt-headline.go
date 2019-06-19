// nytheadlines retrieves data
// from the New York Times API, decodes and displays it.
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
	err = json.NewDecoder(r).Decode(&data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "decode: %v\n", err)
		return
	}
