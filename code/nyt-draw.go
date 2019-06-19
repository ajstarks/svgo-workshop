	// Positioning vars
	thumbsize := 75
	top, left := 200, 150
	imx, titley := left - 100, top-100
	x, y := left, top
	ts := fmt.Sprintf("New York Times for %s", time.Now().Format(datefmt))

	// Draw
	canvas.Gstyle(style)
	canvas.Text(imx, titley, ts, "font-size:175%")
	for _, d := range data.Results {
		canvas.Text(x, y, d.Title)
		canvas.Image(imx, y-thumbsize/2, thumbsize, thumbsize, d.Thumbnail)
		y += 120
	}
	canvas.Gend()
