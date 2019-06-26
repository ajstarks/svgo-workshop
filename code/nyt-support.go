// struncate truncates a string with ellipsis
func struncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[0:n] + "..."
}

// imageinfo returns the thumbnail image information
func imageinfo(data result) (int, int, string) {
	for _, m := range data.Multimedia {
		if m.Format == "Standard Thumbnail" {
			return m.Width, m.Height, m.URL
		}
	}
	return 75, 75, ""
}
