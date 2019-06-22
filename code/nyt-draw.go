	canvas.Gstyle(style)
	canvas.Text(imx, titley, ts, "font-size:150%")
	for i := 0; i < n; i++ {
		d := data.Results[i]
		for _, m := range d.Multimedia {
			if m.Format == "Standard Thumbnail" {
				tw, th, imagelink = m.Width, m.Height, m.URL
				break
			}
		}
		canvas.Text(x, y, d.Title)
		canvas.Image(imx, y-th/2, tw, th, imagelink)
		y += th + (th / 2)
	}
	canvas.Gend()
}
