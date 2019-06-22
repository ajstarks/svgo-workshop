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
	style     = "font-size:20pt;font-family:sans-serif;fill:white"
	errfmt    = "unable to get network data for %s (%s)"
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
	var nstories = flag.Int("n", 5, "number of stories")
	flag.Parse()
	width, height := 1200, *nstories*150
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height)
	nytStories(canvas, *section, *nstories)
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

// drawStories walks the result structure, displaying title and thumbnail
func drawStories(canvas *svg.SVG, data NYTStories, n int) {
	top, left := 200, 150
	imx, titley := left-100, top-100
	x, y, th, tw := left, top, 75, 75
	ts := fmt.Sprintf("New York Times for %s", time.Now().Format(datefmt))
	imagelink := ""
	if n > data.StoryCount {
		n = data.StoryCount
	}

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
