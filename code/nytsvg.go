package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	svg "github.com/ajstarks/svgo"
)

// API Info and formats
const (
	NYTAPIkey = "HxTV50RYD7LyGzHxUy7XCThV3kQ3HYvk"
	NYTfmt    = "https://api.nytimes.com/svc/topstories/v2/%s.json?api-key=%s"
	style     = "font-size:9pt;font-family:sans-serif;text-anchor:middle;fill:white"
	errfmt    = "unable to get network data for %s (%s)"
	headfmt   = "New York Times Top stores for %s"
	datefmt   = "Monday Jan 2, 2006"
	usage     = `section choices:
arts, automobiles, books, business, fashion, 
food, health, home, insider, magazine, movies, 
national, nyregion, obituaries, opinion, politics, 
realestate, science, sports, sundayreview, technology,
theater, tmagazine, travel, upshot, world`
)

// NYTStories is the headline info from the New York Times
type NYTStories struct {
	StoryCount int      `json:"num_results"`
	Results    []result `json:"results"`
}

type result struct {
	Title      string       `json:"title"`
	URL        string       `json:"url"`
	Multimedia []multimedia `json:"multimedia"`
}

type multimedia struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Format string `json:"format"`
}

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

// netread derefernces a URL,
// returning the Reader, with an error
func netread(url string) (io.ReadCloser, error) {
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(errfmt, url, resp.Status)
	}
	return resp.Body, nil
}

// nytStories retrieves Top Stories data
// from the New York Times API, decodes and displays it.
func nytStories(canvas *svg.SVG, width int, section string, n int) {
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
	drawStories(canvas, width, data, n)
}

// drawStories walks the result structure,
// displaying title and thumbnail in a grid
func drawStories(canvas *svg.SVG, width int, data NYTStories, n int) {
	top, left := 100, 150
	titley := top - 50
	x, y := left, top
	ts := fmt.Sprintf(headfmt, time.Now().Format(datefmt))
	if n > data.StoryCount {
		n = data.StoryCount
	}

	canvas.Gstyle(style)
	canvas.Text(width/2, titley, ts, "font-size:300%")
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
