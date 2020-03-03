	canvas.Gstyle(style)
	canvas.Text(width/2, titley, ts, "font-size:250%")
	for i := 0; i < n; i++ {
		d := data.Results[i]
		tw, th, imagelink := imageinfo(d)
		if i > 0 && i%5 == 0 {
			x = left
			y += th + (th / 2)
		}
		canvas.Link(d.URL, d.Title)
		canvas.Image(x, y, tw, th, imagelink)
		canvas.Text(x+tw/2, y+th+12, struncate(d.Title, 20))
		canvas.LinkEnd()
		x += tw * 2
	}
	canvas.Gend()
}
